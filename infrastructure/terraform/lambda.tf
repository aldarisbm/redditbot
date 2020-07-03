resource "aws_lambda_permission" "allow_cloudwatch" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cronjob.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.daily_event.arn
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = data.aws_iam_policy_document.policy_document.json
}

data "aws_iam_policy_document" "policy_document" {
  statement {
    sid     = "AllowAssumeRole"
    actions = ["sts:AssumeRole"]
    effect  = "Allow"
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
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
      SECRET_NAME = aws_secretsmanager_secret.secret.name
      REGION      = var.region
    }
  }
}
