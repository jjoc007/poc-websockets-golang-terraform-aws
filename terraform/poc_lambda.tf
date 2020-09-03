resource "aws_lambda_function" "poc_web_socket_connect_lambda" {
  filename      =  data.archive_file.OnConnectZip.output_path
  function_name = "poc_web_socket_connect_lambda"
  role          = aws_iam_role.websocket_lambda_role.arn
  handler       = "main"
  source_code_hash = data.archive_file.OnConnectZip.output_base64sha256
  runtime = "go1.x"

  environment {
    variables = {
      TABLE_WEB_SOCKET_NAME = aws_dynamodb_table.poc_web_socket_table.name
      API_GATEWAY_ID = aws_apigatewayv2_api.poc_web_socket_api.id
    }
  }
}
resource "aws_lambda_function" "poc_web_socket_disconnect_lambda" {
  filename      = data.archive_file.OnDisconnectZip.output_path
  function_name = "poc_web_socket_disconnect_lambda"
  role          = aws_iam_role.websocket_lambda_role.arn
  handler       = "main"
  source_code_hash = data.archive_file.OnDisconnectZip.output_base64sha256
  runtime = "go1.x"

  environment {
    variables = {
      TABLE_WEB_SOCKET_NAME = aws_dynamodb_table.poc_web_socket_table.name
      API_GATEWAY_ID = aws_apigatewayv2_api.poc_web_socket_api.id
    }
  }
}

resource "aws_lambda_function" "poc_web_socket_message_lambda" {
  filename      = data.archive_file.MessageZip.output_path
  function_name = "poc_web_socket_message_lambda"
  role          = aws_iam_role.websocket_lambda_role.arn
  handler       = "main"
  source_code_hash = data.archive_file.MessageZip.output_base64sha256
  runtime = "go1.x"

  environment {
    variables = {
      TABLE_WEB_SOCKET_NAME = aws_dynamodb_table.poc_web_socket_table.name
      API_GATEWAY_ID = aws_apigatewayv2_api.poc_web_socket_api.id
      ENVIRONMENT = "dev"
    }
  }
}