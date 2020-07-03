resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cronjob.function_name
  principal     = "events.amazonaws.com"
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

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
      "Sid": "AllowAssumeRolegl man"
    }
  ]
}
EOF
}

resource "aws_lambda_function" "cronjob" {
  filename         = "lambda.zip"
  function_name    = "cronjob"
  role             = aws_iam_role.iam_for_lambda.arn
  handler          = "main.handler"
  source_code_hash = filebase64sha256("lambda.zip")
  runtime          = "go1.x"

  environment {
    variables = {
      SECRETS_MANAGER_ARN = aws_secretsmanager_secret.arn
    }
  }
}
