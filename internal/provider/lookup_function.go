package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ function.Function = LookupFunction{}

func NewLookupFunction() function.Function {
	return LookupFunction{}
}

type LookupFunction struct{}

func (f LookupFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookup"
}

func (f LookupFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return a value from a map given its key, or a default value.",
		MarkdownDescription: "Looks up a key in a map. If the key is not present, it returns the provided default value.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "key",
				MarkdownDescription: "The key to look up in the map.",
			},
			function.MapParameter{
				Name:                "input_map",
				MarkdownDescription: "The map to look up the key in.",
				ElementType:         types.StringType,
			},
			function.StringParameter{
				Name:                "default_value",
				MarkdownDescription: "The fallback value if the key is not found.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f LookupFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var key string
	var inputMap map[string]string
	var defaultValue string

	resp.Error = req.Arguments.Get(ctx, &key, &inputMap, &defaultValue)
	if resp.Error != nil {
		return
	}

	val, ok := inputMap[key]
	if ok {
		resp.Error = resp.Result.Set(ctx, val)
	} else {
		resp.Error = resp.Result.Set(ctx, defaultValue)
	}
}
