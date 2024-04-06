
import 'package:citizen_api/citizen_api.dart';

Future<void> main(List<String> args) async {
  final client = CitizenServiceGrpcClient(
    host: 'localhost',
    port: 8080,
  );

  final response = await client.findCitizen('26349413k');
  print('Received: ${response.results}');

  await client.close();
}
