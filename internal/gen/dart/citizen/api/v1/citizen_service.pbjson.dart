//
//  Generated code. Do not modify.
//  source: citizen/api/v1/citizen_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use matchTypeDescriptor instead')
const MatchType$json = {
  '1': 'MatchType',
  '2': [
    {'1': 'MATCH_TYPE_UNSPECIFIED', '2': 0},
    {'1': 'MATCH_TYPE_BY_DOCUMENT_ID', '2': 1},
    {'1': 'MATCH_TYPE_BY_NAME', '2': 2},
  ],
};

/// Descriptor for `MatchType`. Decode as a `google.protobuf.EnumDescriptorProto`.
final $typed_data.Uint8List matchTypeDescriptor = $convert.base64Decode(
    'CglNYXRjaFR5cGUSGgoWTUFUQ0hfVFlQRV9VTlNQRUNJRklFRBAAEh0KGU1BVENIX1RZUEVfQl'
    'lfRE9DVU1FTlRfSUQQARIWChJNQVRDSF9UWVBFX0JZX05BTUUQAg==');

@$core.Deprecated('Use locationDescriptor instead')
const Location$json = {
  '1': 'Location',
  '2': [
    {'1': 'country', '3': 1, '4': 1, '5': 9, '10': 'country'},
    {'1': 'state', '3': 2, '4': 1, '5': 9, '10': 'state'},
    {'1': 'municipality', '3': 3, '4': 1, '5': 9, '10': 'municipality'},
    {'1': 'parish', '3': 4, '4': 1, '5': 9, '10': 'parish'},
    {'1': 'location_id', '3': 5, '4': 1, '5': 9, '10': 'locationId'},
    {'1': 'latitude', '3': 6, '4': 1, '5': 1, '10': 'latitude'},
    {'1': 'longitude', '3': 7, '4': 1, '5': 1, '10': 'longitude'},
  ],
};

/// Descriptor for `Location`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List locationDescriptor = $convert.base64Decode(
    'CghMb2NhdGlvbhIYCgdjb3VudHJ5GAEgASgJUgdjb3VudHJ5EhQKBXN0YXRlGAIgASgJUgVzdG'
    'F0ZRIiCgxtdW5pY2lwYWxpdHkYAyABKAlSDG11bmljaXBhbGl0eRIWCgZwYXJpc2gYBCABKAlS'
    'BnBhcmlzaBIfCgtsb2NhdGlvbl9pZBgFIAEoCVIKbG9jYXRpb25JZBIaCghsYXRpdHVkZRgGIA'
    'EoAVIIbGF0aXR1ZGUSHAoJbG9uZ2l0dWRlGAcgASgBUglsb25naXR1ZGU=');

@$core.Deprecated('Use documentIDDescriptor instead')
const DocumentID$json = {
  '1': 'DocumentID',
  '2': [
    {'1': 'number', '3': 1, '4': 1, '5': 9, '10': 'number'},
    {'1': 'location', '3': 2, '4': 1, '5': 11, '6': '.shipment.api.v1.Location', '10': 'location'},
  ],
};

/// Descriptor for `DocumentID`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List documentIDDescriptor = $convert.base64Decode(
    'CgpEb2N1bWVudElEEhYKBm51bWJlchgBIAEoCVIGbnVtYmVyEjUKCGxvY2F0aW9uGAIgASgLMh'
    'kuc2hpcG1lbnQuYXBpLnYxLkxvY2F0aW9uUghsb2NhdGlvbg==');

@$core.Deprecated('Use citizenDescriptor instead')
const Citizen$json = {
  '1': 'Citizen',
  '2': [
    {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    {'1': 'documents', '3': 2, '4': 3, '5': 11, '6': '.shipment.api.v1.DocumentID', '10': 'documents'},
  ],
};

/// Descriptor for `Citizen`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List citizenDescriptor = $convert.base64Decode(
    'CgdDaXRpemVuEhIKBG5hbWUYASABKAlSBG5hbWUSOQoJZG9jdW1lbnRzGAIgAygLMhsuc2hpcG'
    '1lbnQuYXBpLnYxLkRvY3VtZW50SURSCWRvY3VtZW50cw==');

@$core.Deprecated('Use findCitizenResultDescriptor instead')
const FindCitizenResult$json = {
  '1': 'FindCitizenResult',
  '2': [
    {'1': 'citizen', '3': 1, '4': 1, '5': 11, '6': '.shipment.api.v1.Citizen', '10': 'citizen'},
    {'1': 'match_type', '3': 2, '4': 1, '5': 14, '6': '.shipment.api.v1.MatchType', '10': 'matchType'},
  ],
};

/// Descriptor for `FindCitizenResult`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List findCitizenResultDescriptor = $convert.base64Decode(
    'ChFGaW5kQ2l0aXplblJlc3VsdBIyCgdjaXRpemVuGAEgASgLMhguc2hpcG1lbnQuYXBpLnYxLk'
    'NpdGl6ZW5SB2NpdGl6ZW4SOQoKbWF0Y2hfdHlwZRgCIAEoDjIaLnNoaXBtZW50LmFwaS52MS5N'
    'YXRjaFR5cGVSCW1hdGNoVHlwZQ==');

@$core.Deprecated('Use findCitizenByDocIdRequestDescriptor instead')
const FindCitizenByDocIdRequest$json = {
  '1': 'FindCitizenByDocIdRequest',
  '2': [
    {'1': 'document_id', '3': 1, '4': 1, '5': 9, '10': 'documentId'},
  ],
};

/// Descriptor for `FindCitizenByDocIdRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List findCitizenByDocIdRequestDescriptor = $convert.base64Decode(
    'ChlGaW5kQ2l0aXplbkJ5RG9jSWRSZXF1ZXN0Eh8KC2RvY3VtZW50X2lkGAEgASgJUgpkb2N1bW'
    'VudElk');

@$core.Deprecated('Use findCitizenByDocIdResponseDescriptor instead')
const FindCitizenByDocIdResponse$json = {
  '1': 'FindCitizenByDocIdResponse',
  '2': [
    {'1': 'results', '3': 1, '4': 3, '5': 11, '6': '.shipment.api.v1.FindCitizenResult', '10': 'results'},
  ],
};

/// Descriptor for `FindCitizenByDocIdResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List findCitizenByDocIdResponseDescriptor = $convert.base64Decode(
    'ChpGaW5kQ2l0aXplbkJ5RG9jSWRSZXNwb25zZRI8CgdyZXN1bHRzGAEgAygLMiIuc2hpcG1lbn'
    'QuYXBpLnYxLkZpbmRDaXRpemVuUmVzdWx0UgdyZXN1bHRz');

const $core.Map<$core.String, $core.dynamic> CitizenServiceBase$json = {
  '1': 'CitizenService',
  '2': [
    {'1': 'FindCitizenByDocId', '2': '.shipment.api.v1.FindCitizenByDocIdRequest', '3': '.shipment.api.v1.FindCitizenByDocIdResponse'},
  ],
};

@$core.Deprecated('Use citizenServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> CitizenServiceBase$messageJson = {
  '.shipment.api.v1.FindCitizenByDocIdRequest': FindCitizenByDocIdRequest$json,
  '.shipment.api.v1.FindCitizenByDocIdResponse': FindCitizenByDocIdResponse$json,
  '.shipment.api.v1.FindCitizenResult': FindCitizenResult$json,
  '.shipment.api.v1.Citizen': Citizen$json,
  '.shipment.api.v1.DocumentID': DocumentID$json,
  '.shipment.api.v1.Location': Location$json,
};

/// Descriptor for `CitizenService`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List citizenServiceDescriptor = $convert.base64Decode(
    'Cg5DaXRpemVuU2VydmljZRJtChJGaW5kQ2l0aXplbkJ5RG9jSWQSKi5zaGlwbWVudC5hcGkudj'
    'EuRmluZENpdGl6ZW5CeURvY0lkUmVxdWVzdBorLnNoaXBtZW50LmFwaS52MS5GaW5kQ2l0aXpl'
    'bkJ5RG9jSWRSZXNwb25zZQ==');

