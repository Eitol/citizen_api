//
//  Generated code. Do not modify.
//  source: citizen/api/v1/citizen_service.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class MatchType extends $pb.ProtobufEnum {
  static const MatchType MATCH_TYPE_UNSPECIFIED = MatchType._(0, _omitEnumNames ? '' : 'MATCH_TYPE_UNSPECIFIED');
  static const MatchType MATCH_TYPE_BY_DOCUMENT_ID = MatchType._(1, _omitEnumNames ? '' : 'MATCH_TYPE_BY_DOCUMENT_ID');
  static const MatchType MATCH_TYPE_BY_NAME = MatchType._(2, _omitEnumNames ? '' : 'MATCH_TYPE_BY_NAME');

  static const $core.List<MatchType> values = <MatchType> [
    MATCH_TYPE_UNSPECIFIED,
    MATCH_TYPE_BY_DOCUMENT_ID,
    MATCH_TYPE_BY_NAME,
  ];

  static final $core.Map<$core.int, MatchType> _byValue = $pb.ProtobufEnum.initByValue(values);
  static MatchType? valueOf($core.int value) => _byValue[value];

  const MatchType._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
