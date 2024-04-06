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

import 'package:protobuf/protobuf.dart' as $pb;

import 'citizen_service.pbenum.dart';

export 'citizen_service.pbenum.dart';

class Location extends $pb.GeneratedMessage {
  factory Location({
    $core.String? country,
    $core.String? state,
    $core.String? municipality,
    $core.String? parish,
    $core.String? locationId,
    $core.double? latitude,
    $core.double? longitude,
  }) {
    final $result = create();
    if (country != null) {
      $result.country = country;
    }
    if (state != null) {
      $result.state = state;
    }
    if (municipality != null) {
      $result.municipality = municipality;
    }
    if (parish != null) {
      $result.parish = parish;
    }
    if (locationId != null) {
      $result.locationId = locationId;
    }
    if (latitude != null) {
      $result.latitude = latitude;
    }
    if (longitude != null) {
      $result.longitude = longitude;
    }
    return $result;
  }
  Location._() : super();
  factory Location.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Location.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Location', package: const $pb.PackageName(_omitMessageNames ? '' : 'shipment.api.v1'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'country')
    ..aOS(2, _omitFieldNames ? '' : 'state')
    ..aOS(3, _omitFieldNames ? '' : 'municipality')
    ..aOS(4, _omitFieldNames ? '' : 'parish')
    ..aOS(5, _omitFieldNames ? '' : 'locationId')
    ..a<$core.double>(6, _omitFieldNames ? '' : 'latitude', $pb.PbFieldType.OD)
    ..a<$core.double>(7, _omitFieldNames ? '' : 'longitude', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Location clone() => Location()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Location copyWith(void Function(Location) updates) => super.copyWith((message) => updates(message as Location)) as Location;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Location create() => Location._();
  Location createEmptyInstance() => create();
  static $pb.PbList<Location> createRepeated() => $pb.PbList<Location>();
  @$core.pragma('dart2js:noInline')
  static Location getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Location>(create);
  static Location? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get country => $_getSZ(0);
  @$pb.TagNumber(1)
  set country($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCountry() => $_has(0);
  @$pb.TagNumber(1)
  void clearCountry() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get state => $_getSZ(1);
  @$pb.TagNumber(2)
  set state($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasState() => $_has(1);
  @$pb.TagNumber(2)
  void clearState() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get municipality => $_getSZ(2);
  @$pb.TagNumber(3)
  set municipality($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMunicipality() => $_has(2);
  @$pb.TagNumber(3)
  void clearMunicipality() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get parish => $_getSZ(3);
  @$pb.TagNumber(4)
  set parish($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasParish() => $_has(3);
  @$pb.TagNumber(4)
  void clearParish() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get locationId => $_getSZ(4);
  @$pb.TagNumber(5)
  set locationId($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLocationId() => $_has(4);
  @$pb.TagNumber(5)
  void clearLocationId() => clearField(5);

  @$pb.TagNumber(6)
  $core.double get latitude => $_getN(5);
  @$pb.TagNumber(6)
  set latitude($core.double v) { $_setDouble(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasLatitude() => $_has(5);
  @$pb.TagNumber(6)
  void clearLatitude() => clearField(6);

  @$pb.TagNumber(7)
  $core.double get longitude => $_getN(6);
  @$pb.TagNumber(7)
  set longitude($core.double v) { $_setDouble(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasLongitude() => $_has(6);
  @$pb.TagNumber(7)
  void clearLongitude() => clearField(7);
}

class DocumentID extends $pb.GeneratedMessage {
  factory DocumentID({
    $core.String? number,
    Location? location,
  }) {
    final $result = create();
    if (number != null) {
      $result.number = number;
    }
    if (location != null) {
      $result.location = location;
    }
    return $result;
  }
  DocumentID._() : super();
  factory DocumentID.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory DocumentID.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'DocumentID', package: const $pb.PackageName(_omitMessageNames ? '' : 'shipment.api.v1'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'number')
    ..aOM<Location>(2, _omitFieldNames ? '' : 'location', subBuilder: Location.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  DocumentID clone() => DocumentID()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  DocumentID copyWith(void Function(DocumentID) updates) => super.copyWith((message) => updates(message as DocumentID)) as DocumentID;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static DocumentID create() => DocumentID._();
  DocumentID createEmptyInstance() => create();
  static $pb.PbList<DocumentID> createRepeated() => $pb.PbList<DocumentID>();
  @$core.pragma('dart2js:noInline')
  static DocumentID getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<DocumentID>(create);
  static DocumentID? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get number => $_getSZ(0);
  @$pb.TagNumber(1)
  set number($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearNumber() => clearField(1);

  @$pb.TagNumber(2)
  Location get location => $_getN(1);
  @$pb.TagNumber(2)
  set location(Location v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasLocation() => $_has(1);
  @$pb.TagNumber(2)
  void clearLocation() => clearField(2);
  @$pb.TagNumber(2)
  Location ensureLocation() => $_ensure(1);
}

class Citizen extends $pb.GeneratedMessage {
  factory Citizen({
    $core.String? name,
    $core.Iterable<DocumentID>? documents,
  }) {
    final $result = create();
    if (name != null) {
      $result.name = name;
    }
    if (documents != null) {
      $result.documents.addAll(documents);
    }
    return $result;
  }
  Citizen._() : super();
  factory Citizen.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Citizen.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Citizen', package: const $pb.PackageName(_omitMessageNames ? '' : 'shipment.api.v1'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'name')
    ..pc<DocumentID>(2, _omitFieldNames ? '' : 'documents', $pb.PbFieldType.PM, subBuilder: DocumentID.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Citizen clone() => Citizen()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Citizen copyWith(void Function(Citizen) updates) => super.copyWith((message) => updates(message as Citizen)) as Citizen;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Citizen create() => Citizen._();
  Citizen createEmptyInstance() => create();
  static $pb.PbList<Citizen> createRepeated() => $pb.PbList<Citizen>();
  @$core.pragma('dart2js:noInline')
  static Citizen getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Citizen>(create);
  static Citizen? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.List<DocumentID> get documents => $_getList(1);
}

class FindCitizenResult extends $pb.GeneratedMessage {
  factory FindCitizenResult({
    Citizen? citizen,
    MatchType? matchType,
  }) {
    final $result = create();
    if (citizen != null) {
      $result.citizen = citizen;
    }
    if (matchType != null) {
      $result.matchType = matchType;
    }
    return $result;
  }
  FindCitizenResult._() : super();
  factory FindCitizenResult.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FindCitizenResult.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'FindCitizenResult', package: const $pb.PackageName(_omitMessageNames ? '' : 'shipment.api.v1'), createEmptyInstance: create)
    ..aOM<Citizen>(1, _omitFieldNames ? '' : 'citizen', subBuilder: Citizen.create)
    ..e<MatchType>(2, _omitFieldNames ? '' : 'matchType', $pb.PbFieldType.OE, defaultOrMaker: MatchType.MATCH_TYPE_UNSPECIFIED, valueOf: MatchType.valueOf, enumValues: MatchType.values)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FindCitizenResult clone() => FindCitizenResult()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FindCitizenResult copyWith(void Function(FindCitizenResult) updates) => super.copyWith((message) => updates(message as FindCitizenResult)) as FindCitizenResult;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static FindCitizenResult create() => FindCitizenResult._();
  FindCitizenResult createEmptyInstance() => create();
  static $pb.PbList<FindCitizenResult> createRepeated() => $pb.PbList<FindCitizenResult>();
  @$core.pragma('dart2js:noInline')
  static FindCitizenResult getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FindCitizenResult>(create);
  static FindCitizenResult? _defaultInstance;

  @$pb.TagNumber(1)
  Citizen get citizen => $_getN(0);
  @$pb.TagNumber(1)
  set citizen(Citizen v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCitizen() => $_has(0);
  @$pb.TagNumber(1)
  void clearCitizen() => clearField(1);
  @$pb.TagNumber(1)
  Citizen ensureCitizen() => $_ensure(0);

  @$pb.TagNumber(2)
  MatchType get matchType => $_getN(1);
  @$pb.TagNumber(2)
  set matchType(MatchType v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasMatchType() => $_has(1);
  @$pb.TagNumber(2)
  void clearMatchType() => clearField(2);
}

class FindCitizenByDocIdRequest extends $pb.GeneratedMessage {
  factory FindCitizenByDocIdRequest({
    $core.String? documentId,
  }) {
    final $result = create();
    if (documentId != null) {
      $result.documentId = documentId;
    }
    return $result;
  }
  FindCitizenByDocIdRequest._() : super();
  factory FindCitizenByDocIdRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FindCitizenByDocIdRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'FindCitizenByDocIdRequest', package: const $pb.PackageName(_omitMessageNames ? '' : 'shipment.api.v1'), createEmptyInstance: create)
    ..aOS(1, _omitFieldNames ? '' : 'documentId')
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FindCitizenByDocIdRequest clone() => FindCitizenByDocIdRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FindCitizenByDocIdRequest copyWith(void Function(FindCitizenByDocIdRequest) updates) => super.copyWith((message) => updates(message as FindCitizenByDocIdRequest)) as FindCitizenByDocIdRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static FindCitizenByDocIdRequest create() => FindCitizenByDocIdRequest._();
  FindCitizenByDocIdRequest createEmptyInstance() => create();
  static $pb.PbList<FindCitizenByDocIdRequest> createRepeated() => $pb.PbList<FindCitizenByDocIdRequest>();
  @$core.pragma('dart2js:noInline')
  static FindCitizenByDocIdRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FindCitizenByDocIdRequest>(create);
  static FindCitizenByDocIdRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get documentId => $_getSZ(0);
  @$pb.TagNumber(1)
  set documentId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDocumentId() => $_has(0);
  @$pb.TagNumber(1)
  void clearDocumentId() => clearField(1);
}

class FindCitizenByDocIdResponse extends $pb.GeneratedMessage {
  factory FindCitizenByDocIdResponse({
    $core.Iterable<FindCitizenResult>? results,
  }) {
    final $result = create();
    if (results != null) {
      $result.results.addAll(results);
    }
    return $result;
  }
  FindCitizenByDocIdResponse._() : super();
  factory FindCitizenByDocIdResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FindCitizenByDocIdResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'FindCitizenByDocIdResponse', package: const $pb.PackageName(_omitMessageNames ? '' : 'shipment.api.v1'), createEmptyInstance: create)
    ..pc<FindCitizenResult>(1, _omitFieldNames ? '' : 'results', $pb.PbFieldType.PM, subBuilder: FindCitizenResult.create)
    ..hasRequiredFields = false
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  FindCitizenByDocIdResponse clone() => FindCitizenByDocIdResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  FindCitizenByDocIdResponse copyWith(void Function(FindCitizenByDocIdResponse) updates) => super.copyWith((message) => updates(message as FindCitizenByDocIdResponse)) as FindCitizenByDocIdResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static FindCitizenByDocIdResponse create() => FindCitizenByDocIdResponse._();
  FindCitizenByDocIdResponse createEmptyInstance() => create();
  static $pb.PbList<FindCitizenByDocIdResponse> createRepeated() => $pb.PbList<FindCitizenByDocIdResponse>();
  @$core.pragma('dart2js:noInline')
  static FindCitizenByDocIdResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FindCitizenByDocIdResponse>(create);
  static FindCitizenByDocIdResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<FindCitizenResult> get results => $_getList(0);
}

class CitizenServiceApi {
  $pb.RpcClient _client;
  CitizenServiceApi(this._client);

  $async.Future<FindCitizenByDocIdResponse> findCitizenByDocId($pb.ClientContext? ctx, FindCitizenByDocIdRequest request) =>
    _client.invoke<FindCitizenByDocIdResponse>(ctx, 'CitizenService', 'FindCitizenByDocId', request, FindCitizenByDocIdResponse())
  ;
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
