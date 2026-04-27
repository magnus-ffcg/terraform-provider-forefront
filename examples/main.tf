terraform {
  required_providers {
    forefront = {
      source = "magnus-ffcg/terraform-provider-forefront"
    }
  }
}

provider "forefront" {}

output "lookup_example" {
  value = provider::forefront::lookup("my_key", { "my_key" = "Success!" }, "default_value")
}

output "replace_deep_example" {
  value = provider::forefront::replace_deep({
    "nested_list" = [
      { "sub_key" = "foo is bar" }
    ],
    "key2" = "another foo thing"
  }, "foo", "baz")
}
