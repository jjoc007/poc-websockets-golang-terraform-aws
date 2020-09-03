data "archive_file" "OnConnectZip" {
  type        = "zip"
  source_dir = "lambdas/dist/poc_web_socket_connect_lambda"
  output_path = "lambdas/dist/poc_web_socket_connect_lambda.zip"
}

data "archive_file" "OnDisconnectZip" {
  type        = "zip"
  source_dir = "lambdas/dist/poc_web_socket_disconnect_lambda"
  output_path = "lambdas/dist/poc_web_socket_disconnect_lambda.zip"
}

data "archive_file" "MessageZip" {
  type        = "zip"
  source_dir = "lambdas/dist/poc_web_socket_messages_lambda"
  output_path = "lambdas/dist/poc_web_socket_messages_lambda.zip"
}