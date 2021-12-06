locals {
  functions = yamldecode(file(".sls-tf/functions.yml"))
}
