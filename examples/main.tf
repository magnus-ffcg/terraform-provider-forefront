terraform {
  required_providers {
    forefront = {
      source = "magnus-ffcg/forefront"
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

output "replace_deep_map_example" {
  value = provider::forefront::replace_deep_map({
    "nested_list" = [
      { "sub_key" = "foo is bar" }
    ],
    "key2" = "another foo thing and world"
    }, {
    "foo"   = "baz"
    "world" = "earth"
    "bar"   = "qux"
  })
}

