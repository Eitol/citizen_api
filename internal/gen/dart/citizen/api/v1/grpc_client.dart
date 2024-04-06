import 'package:grpc/grpc.dart';

import 'citizen_service.pbgrpc.dart';


class CitizenServiceGrpcClient {
  late ClientChannel _channel;
  late CitizenServiceClient _client;

  CitizenServiceGrpcClient({required String host, required int port}) {
    _channel = ClientChannel(
      host,
      port: port,
      options: ChannelOptions(
        credentials: ChannelCredentials.insecure(),
      ),
    );
    _client = CitizenServiceClient(_channel);
  }

  Future<FindCitizenByDocIdResponse> findCitizen(String documentId) async {
    var request = FindCitizenByDocIdRequest(documentId: documentId);
    return _client.findCitizenByDocId(request);
  }

  Future<void> close() async {
    await _channel.shutdown();
  }
}