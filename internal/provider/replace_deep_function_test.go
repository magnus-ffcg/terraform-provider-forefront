package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccReplaceDeepFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "replaced" {
				  value = jsonencode(provider::forefront::replace_deep({
                      "nested_list" = [
                          { "sub_key" = "foo is bar" }
                      ],
                      "key2" = "another foo thing"
                  }, "foo", "baz"))
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("replaced", `{"key2":"another baz thing","nested_list":[{"sub_key":"baz is bar"}]}`),
				),
			},
		},
	})
}
