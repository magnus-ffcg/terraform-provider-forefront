package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLookupFunction(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "found" {
				  value = provider::forefront::lookup("foo", { "foo" = "bar" }, "default")
				}

				output "not_found" {
				  value = provider::forefront::lookup("missing", { "foo" = "bar" }, "default")
				}
				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckOutput("found", "bar"),
					resource.TestCheckOutput("not_found", "default"),
				),
			},
		},
	})
}
