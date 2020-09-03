resource "aws_apigatewayv2_api" "poc_web_socket_api" {
  name                       = "poc_web_socket_api"
  protocol_type              = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_deployment" "Deployment" {
  api_id = aws_apigatewayv2_api.poc_web_socket_api.id

  depends_on = [
    aws_apigatewayv2_route.connect_route,
    aws_apigatewayv2_route.disconnect_route,
    aws_apigatewayv2_route.message_route,
  ]
}

resource "aws_apigatewayv2_stage" "poc_web_socket_api_stage" {
  api_id        = aws_apigatewayv2_api.poc_web_socket_api.id
  name          = "dev"
  deployment_id = aws_apigatewayv2_deployment.Deployment.id
}


resource "aws_apigatewayv2_integration" "poc_ws_connect_integration" {
  api_id             = aws_apigatewayv2_api.poc_web_socket_api.id
  integration_type   = "AWS_PROXY"
  description        = "Lambda example"
  integration_uri    = aws_lambda_function.poc_web_socket_connect_lambda.invoke_arn
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "connect_route" {
  api_id         = aws_apigatewayv2_api.poc_web_socket_api.id
  route_key      = "$connect"
  operation_name = "ConnectRoute"
  target         = "integrations/${aws_apigatewayv2_integration.poc_ws_connect_integration.id}"
  authorization_type = "NONE"
}

resource "aws_apigatewayv2_integration" "poc_ws_disconnect_integration" {
  api_id             = aws_apigatewayv2_api.poc_web_socket_api.id
  integration_type   = "AWS_PROXY"
  description        = "Lambda example"
  integration_uri    = aws_lambda_function.poc_web_socket_disconnect_lambda.invoke_arn
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "disconnect_route" {
  api_id    = aws_apigatewayv2_api.poc_web_socket_api.id
  route_key = "$disconnect"
  operation_name = "DisconnectRoute"
  target         = "integrations/${aws_apigatewayv2_integration.poc_ws_disconnect_integration.id}"
}

resource "aws_apigatewayv2_integration" "poc_ws_message_integration" {
  api_id             = aws_apigatewayv2_api.poc_web_socket_api.id
  integration_type   = "AWS_PROXY"
  description        = "Lambda example"
  integration_uri    = aws_lambda_function.poc_web_socket_message_lambda.invoke_arn
  integration_method = "POST"
}

resource "aws_apigatewayv2_route" "message_route" {
  api_id    = aws_apigatewayv2_api.poc_web_socket_api.id
  route_key = "message"
  operation_name = "MessageRoute"
  target         = "integrations/${aws_apigatewayv2_integration.poc_ws_message_integration.id}"
}