resource "aws_cloudwatch_event_rule" "daily_event" {
  description         = "Triggers lambda to handle events based on the day of the week"
  schedule_expression = "cron(0 12 ? * MON-FRI *)"
}

resource "aws_cloudwatch_event_target" "lambda" {
  rule      = aws_cloudwatch_event_rule.daily_event.name
  target_id = "SendToLinuxChallengeLambda"
  arn       = aws_lambda_function.cronjob.arn
}