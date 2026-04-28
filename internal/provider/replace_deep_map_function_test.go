package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccReplaceDeepMapFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "replaced" {
				  value = jsonencode(provider::forefront::replace_deep_map({
                      "nested_list" = [
                          { "sub_key" = "foo is bar" }
                      ],
                      "key2" = "another foo thing and world"
                  }, {
					 "foo" = "baz"
					 "world" = "earth"
					 "bar" = "qux"
				  }))
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("replaced", `{"key2":"another baz thing and earth","nested_list":[{"sub_key":"baz is qux"}]}`),
				),
			},
		},
	})
}
