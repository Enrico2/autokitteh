// @generated by protoc-gen-connect-es v1.1.4 with parameter "target=ts"
// @generated from file autokitteh/apply/v1/svc.proto (package autokitteh.apply.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { ApplyRequest, ApplyResponse } from "./svc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service autokitteh.apply.v1.ApplyService
 */
export const ApplyService = {
  typeName: "autokitteh.apply.v1.ApplyService",
  methods: {
    /**
     * @generated from rpc autokitteh.apply.v1.ApplyService.Apply
     */
    apply: {
      name: "Apply",
      I: ApplyRequest,
      O: ApplyResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

