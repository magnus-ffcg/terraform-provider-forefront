package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ function.Function = ReplaceDeepMapFunction{}

func NewReplaceDeepMapFunction() function.Function {
	return ReplaceDeepMapFunction{}
}

type ReplaceDeepMapFunction struct{}

func (f ReplaceDeepMapFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "replace_deep_map"
}

func (f ReplaceDeepMapFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Search and replace strings deeply within any object using a lookup map.",
		MarkdownDescription: "Recursively iterates through lists, maps, sets, tuples, and objects to find string values. Replaces multiple substrings based on a lookup map of find-replace pairs.",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name:                "input",
				MarkdownDescription: "The complex object or collection to process.",
			},
			function.MapParameter{
				Name:                "lookup_map",
				MarkdownDescription: "A map where keys are search strings and values are replace strings.",
				ElementType:         types.StringType,
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (f ReplaceDeepMapFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inputVal basetypes.DynamicValue
	var lookupMap map[string]string

	resp.Error = req.Arguments.Get(ctx, &inputVal, &lookupMap)
	if resp.Error != nil {
		return
	}

	searchReplacePairs := make([]string, 0, len(lookupMap)*2)
	for k, v := range lookupMap {
		searchReplacePairs = append(searchReplacePairs, k, v)
	}
	replacer := strings.NewReplacer(searchReplacePairs...)

	newUnder, err := replaceDeepVal(ctx, inputVal, func(s string) string {
		return replacer.Replace(s)
	})
	if err != nil {
		resp.Error = function.NewFuncError(err.Error())
		return
	}

	resultDynamic, ok := newUnder.(basetypes.DynamicValue)
	if !ok {
		resultDynamic = basetypes.NewDynamicValue(newUnder)
	}

	resp.Error = resp.Result.Set(ctx, resultDynamic)
}
