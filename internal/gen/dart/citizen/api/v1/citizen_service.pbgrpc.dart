//
//  Generated code. Do not modify.
//  source: citizen/api/v1/citizen_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'citizen_service.pb.dart' as $0;

export 'citizen_service.pb.dart';

@$pb.GrpcServiceName('shipment.api.v1.CitizenService')
class CitizenServiceClient extends $grpc.Client {
  static final _$findCitizenByDocId = $grpc.ClientMethod<$0.FindCitizenByDocIdRequest, $0.FindCitizenByDocIdResponse>(
      '/shipment.api.v1.CitizenService/FindCitizenByDocId',
      ($0.FindCitizenByDocIdRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.FindCitizenByDocIdResponse.fromBuffer(value));

  CitizenServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.FindCitizenByDocIdResponse> findCitizenByDocId($0.FindCitizenByDocIdRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$findCitizenByDocId, request, options: options);
  }
}

@$pb.GrpcServiceName('shipment.api.v1.CitizenService')
abstract class CitizenServiceBase extends $grpc.Service {
  $core.String get $name => 'shipment.api.v1.CitizenService';

  CitizenServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.FindCitizenByDocIdRequest, $0.FindCitizenByDocIdResponse>(
        'FindCitizenByDocId',
        findCitizenByDocId_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.FindCitizenByDocIdRequest.fromBuffer(value),
        ($0.FindCitizenByDocIdResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.FindCitizenByDocIdResponse> findCitizenByDocId_Pre($grpc.ServiceCall call, $async.Future<$0.FindCitizenByDocIdRequest> request) async {
    return findCitizenByDocId(call, await request);
  }

  $async.Future<$0.FindCitizenByDocIdResponse> findCitizenByDocId($grpc.ServiceCall call, $0.FindCitizenByDocIdRequest request);
}
