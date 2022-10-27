terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }
}

provider "aws" {
  region = "us-west-2"
  shared_credentials_files = ["$HOME/.aws/credentials"]
}


resource "aws_dynamodb_table" "game-score-table" {
  name = "game-score-table"
  billing_mode = "PROVISIONED"
  read_capacity= "5"
  write_capacity= "5"

  attribute {
    name = "scoreId"
    type = "S"
  }

  hash_key = "scoreId"
}
