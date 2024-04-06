//
//  Generated code. Do not modify.
//  source: citizen/api/v1/citizen_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names
// ignore_for_file: deprecated_member_use_from_same_package, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'citizen_service.pb.dart' as $0;
import 'citizen_service.pbjson.dart';

export 'citizen_service.pb.dart';

abstract class CitizenServiceBase extends $pb.GeneratedService {
  $async.Future<$0.FindCitizenByDocIdResponse> findCitizenByDocId($pb.ServerContext ctx, $0.FindCitizenByDocIdRequest request);

  $pb.GeneratedMessage createRequest($core.String methodName) {
    switch (methodName) {
      case 'FindCitizenByDocId': return $0.FindCitizenByDocIdRequest();
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String methodName, $pb.GeneratedMessage request) {
    switch (methodName) {
      case 'FindCitizenByDocId': return this.findCitizenByDocId(ctx, request as $0.FindCitizenByDocIdRequest);
      default: throw $core.ArgumentError('Unknown method: $methodName');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => CitizenServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => CitizenServiceBase$messageJson;
}

