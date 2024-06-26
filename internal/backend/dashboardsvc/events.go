package dashboardsvc

import (
	"net/http"
	"strconv"

	"go.autokitteh.dev/autokitteh/internal/kittehs"
	"go.autokitteh.dev/autokitteh/sdk/sdkservices"
	"go.autokitteh.dev/autokitteh/sdk/sdktypes"
)

func (s Svc) initEvents() {
	s.Muxes.Auth.HandleFunc("/events", s.events)
	s.Muxes.Auth.HandleFunc("/events/{eid}", s.event)
}

type event struct{ sdktypes.Event }

func (p event) FieldsOrder() []string {
	return []string{"event_id", "connection_id", "integration_id"}
}

func (p event) HideFields() []string { return nil }

func toEvent(sdkP sdktypes.Event) event { return event{sdkP} }

func (s Svc) listEvents(w http.ResponseWriter, r *http.Request, f sdkservices.ListEventsFilter) (list, error) {
	if f.Limit <= 0 {
		f.Limit = 50

		if l := r.URL.Query().Get("max_events"); l != "" {
			l64, err := strconv.ParseInt(l, 10, 32)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return list{}, err

			}

			f.Limit = int(l64)
		}
	}

	if f.MinSequenceNumber == 0 {
		if l := r.URL.Query().Get("min_event_seq"); l != "" {
			l64, err := strconv.ParseInt(l, 10, 32)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return list{}, err

			}

			f.MinSequenceNumber = uint64(l64)
		}
	}

	sdkCs, err := s.Svcs.Events().List(r.Context(), f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return list{}, err
	}

	var drops []string
	if f.IntegrationID.IsValid() {
		drops = append(drops, "integration_id")
	}

	if f.ConnectionID.IsValid() {
		drops = append(drops, "connection_id")
	}

	return genListData(kittehs.Transform(sdkCs, toEvent), drops...), nil
}

func (s Svc) events(w http.ResponseWriter, r *http.Request) {
	ts, err := s.listEvents(w, r, sdkservices.ListEventsFilter{})
	if err != nil {
		return
	}

	renderList(w, r, "events", ts)
}

func (s Svc) event(w http.ResponseWriter, r *http.Request) {
	eid, err := sdktypes.ParseEventID(r.PathValue("eid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sdkP, err := s.Svcs.Events().Get(r.Context(), eid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p := toEvent(sdkP)

	renderBigObject(w, r, "event", p.ToProto())
}
