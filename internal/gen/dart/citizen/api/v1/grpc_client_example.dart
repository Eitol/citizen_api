import 'grpc_client.dart';

Future<void> main(List<String> args) async {
  final client = CitizenServiceGrpcClient(
    host: 'localhost',
    port: 8080,
  );

  final response = await client.findCitizen('26349413k');
  print('Received: ${response.results}');

  await client.close();
}
