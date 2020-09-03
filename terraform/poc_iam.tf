resource "aws_iam_role" "websocket_lambda_role" {
  name               = "websocket_messages_lambda_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "dynamo_policy_web_sockets_messages" {
  name   = "websockets_messages_dynamo_policy_all"
  policy = data.aws_iam_policy_document.dynamo_access_policy.json
}

resource "aws_iam_policy" "default_policy_web_sockets_messages" {
  name   = "websockets_messages_default_policy_all"
  policy = data.aws_iam_policy_document.default_lambda_policy.json
}

resource "aws_iam_policy" "execute_api_policy_web_sockets_messages" {
  name   = "websockets_messages_execute_api_policy_all"
  policy = data.aws_iam_policy_document.execute_api_access_policy.json
}

resource "aws_iam_role_policy_attachment" "lambda_ws_to_dynamo" {
  role       = aws_iam_role.websocket_lambda_role.name
  policy_arn = aws_iam_policy.dynamo_policy_web_sockets_messages.arn
}

resource "aws_iam_role_policy_attachment" "lambda_ws_to_default" {
  role       = aws_iam_role.websocket_lambda_role.name
  policy_arn = aws_iam_policy.default_policy_web_sockets_messages.arn
}

resource "aws_iam_role_policy_attachment" "lambda_ws_to_execute_api" {
  role       = aws_iam_role.websocket_lambda_role.name
  policy_arn = aws_iam_policy.execute_api_policy_web_sockets_messages.arn
}