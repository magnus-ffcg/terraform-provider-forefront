package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ function.Function = ReplaceDeepFunction{}

func NewReplaceDeepFunction() function.Function {
	return ReplaceDeepFunction{}
}

type ReplaceDeepFunction struct{}

func (f ReplaceDeepFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "replace_deep"
}

func (f ReplaceDeepFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Search and replace strings deeply within any object.",
		MarkdownDescription: "Recursively iterates through lists, maps, sets, tuples, and objects to find and replace string values.",
		Parameters: []function.Parameter{
			function.DynamicParameter{
				Name:                "input",
				MarkdownDescription: "The complex object or collection to process.",
			},
			function.StringParameter{
				Name:                "search",
				MarkdownDescription: "The string to search for.",
			},
			function.StringParameter{
				Name:                "replace",
				MarkdownDescription: "The string to replace the search string with.",
			},
		},
		Return: function.DynamicReturn{},
	}
}

func (f ReplaceDeepFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var inputVal basetypes.DynamicValue
	var search string
	var replace string

	resp.Error = req.Arguments.Get(ctx, &inputVal, &search, &replace)
	if resp.Error != nil {
		return
	}

	newUnder, err := replaceDeepVal(ctx, inputVal, func(s string) string {
		return strings.ReplaceAll(s, search, replace)
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

func replaceDeepVal(ctx context.Context, v attr.Value, modifier func(string) string) (attr.Value, error) {
	if v == nil {
		return nil, nil
	}

	switch val := v.(type) {
	case basetypes.DynamicValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		newUnder, err := replaceDeepVal(ctx, val.UnderlyingValue(), modifier)
		if err != nil {
			return val, err
		}
		return basetypes.NewDynamicValue(newUnder), nil
	case basetypes.StringValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		return basetypes.NewStringValue(modifier(val.ValueString())), nil
	case basetypes.ObjectValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		attrs := val.Attributes()
		newAttrs := make(map[string]attr.Value, len(attrs))
		for k, attrV := range attrs {
			newAttrV, err := replaceDeepVal(ctx, attrV, modifier)
			if err != nil {
				return val, err
			}
			newAttrs[k] = newAttrV
		}
		res, diags := basetypes.NewObjectValue(val.AttributeTypes(ctx), newAttrs)
		if diags.HasError() {
			// fallback mapping
		}
		return res, nil
	case basetypes.MapValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		elems := val.Elements()
		newElems := make(map[string]attr.Value, len(elems))
		for k, elemV := range elems {
			newElemV, err := replaceDeepVal(ctx, elemV, modifier)
			if err != nil {
				return val, err
			}
			newElems[k] = newElemV
		}
		res, _ := basetypes.NewMapValue(val.ElementType(ctx), newElems)
		return res, nil

	case basetypes.ListValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		elems := val.Elements()
		newElems := make([]attr.Value, len(elems))
		for i, elemV := range elems {
			newElemV, err := replaceDeepVal(ctx, elemV, modifier)
			if err != nil {
				return val, err
			}
			newElems[i] = newElemV
		}
		res, _ := basetypes.NewListValue(val.ElementType(ctx), newElems)
		return res, nil

	case basetypes.SetValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		elems := val.Elements()
		newElems := make([]attr.Value, len(elems))
		for i, elemV := range elems {
			newElemV, err := replaceDeepVal(ctx, elemV, modifier)
			if err != nil {
				return val, err
			}
			newElems[i] = newElemV
		}
		res, _ := basetypes.NewSetValue(val.ElementType(ctx), newElems)
		return res, nil

	case basetypes.TupleValue:
		if val.IsNull() || val.IsUnknown() {
			return val, nil
		}
		elems := val.Elements()
		newElems := make([]attr.Value, len(elems))
		for i, elemV := range elems {
			newElemV, err := replaceDeepVal(ctx, elemV, modifier)
			if err != nil {
				return val, err
			}
			newElems[i] = newElemV
		}
		res, _ := basetypes.NewTupleValue(val.ElementTypes(ctx), newElems)
		return res, nil

	default:
		// Not a collection or string (e.g. Int64Value, Float64Value, BoolValue)
		return v, nil
	}
}
