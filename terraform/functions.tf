resource "aws_lambda_function" "function" {
  for_each = { for function in local.functions : function.name => function }

  name     = each.key
  filename = each.value.filename
  role     = each.value.role
  handler  = each.value.handler
  runtime  = each.value.runtime
}
