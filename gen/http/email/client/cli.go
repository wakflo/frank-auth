// Code generated by goa v3.20.0, DO NOT EDIT.
//
// email HTTP client CLI support package
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	email "github.com/juicycleff/frank/gen/email"
	goa "goa.design/goa/v3/pkg"
)

// BuildListTemplatesPayload builds the payload for the email list_templates
// endpoint from CLI flags.
func BuildListTemplatesPayload(emailListTemplatesOffset string, emailListTemplatesLimit string, emailListTemplatesType string, emailListTemplatesOrganizationID string, emailListTemplatesLocale string, emailListTemplatesJWT string) (*email.ListTemplatesPayload, error) {
	var err error
	var offset int
	{
		if emailListTemplatesOffset != "" {
			var v int64
			v, err = strconv.ParseInt(emailListTemplatesOffset, 10, strconv.IntSize)
			offset = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for offset, must be INT")
			}
			if offset < 0 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("offset", offset, 0, true))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var limit int
	{
		if emailListTemplatesLimit != "" {
			var v int64
			v, err = strconv.ParseInt(emailListTemplatesLimit, 10, strconv.IntSize)
			limit = int(v)
			if err != nil {
				return nil, fmt.Errorf("invalid value for limit, must be INT")
			}
			if limit < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 1, true))
			}
			if limit > 100 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 100, false))
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var type_ *string
	{
		if emailListTemplatesType != "" {
			type_ = &emailListTemplatesType
		}
	}
	var organizationID *string
	{
		if emailListTemplatesOrganizationID != "" {
			organizationID = &emailListTemplatesOrganizationID
		}
	}
	var locale *string
	{
		if emailListTemplatesLocale != "" {
			locale = &emailListTemplatesLocale
		}
	}
	var jwt *string
	{
		if emailListTemplatesJWT != "" {
			jwt = &emailListTemplatesJWT
		}
	}
	v := &email.ListTemplatesPayload{}
	v.Offset = offset
	v.Limit = limit
	v.Type = type_
	v.OrganizationID = organizationID
	v.Locale = locale
	v.JWT = jwt

	return v, nil
}

// BuildCreateTemplatePayload builds the payload for the email create_template
// endpoint from CLI flags.
func BuildCreateTemplatePayload(emailCreateTemplateBody string, emailCreateTemplateJWT string) (*email.CreateTemplatePayload, error) {
	var err error
	var body CreateTemplateRequestBody
	{
		err = json.Unmarshal([]byte(emailCreateTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"active\": true,\n      \"html_content\": \"\\u003chtml\\u003e\\u003cbody\\u003e\\u003ch1\\u003eWelcome!\\u003c/h1\\u003e\\u003cp\\u003eHello {{name}}\\u003c/p\\u003e\\u003c/body\\u003e\\u003c/html\\u003e\",\n      \"locale\": \"Deserunt autem sit tempora temporibus provident hic.\",\n      \"metadata\": {\n         \"Dicta dignissimos eum est.\": \"Recusandae quam.\"\n      },\n      \"name\": \"Welcome Email\",\n      \"organization_id\": \"Possimus quasi dignissimos ad quisquam et repellat.\",\n      \"subject\": \"Welcome to our platform\",\n      \"system\": false,\n      \"text_content\": \"Welcome! Hello {{name}}\",\n      \"type\": \"welcome\"\n   }'")
		}
	}
	var jwt *string
	{
		if emailCreateTemplateJWT != "" {
			jwt = &emailCreateTemplateJWT
		}
	}
	v := &email.CreateTemplatePayload{
		Name:           body.Name,
		Subject:        body.Subject,
		Type:           body.Type,
		HTMLContent:    body.HTMLContent,
		TextContent:    body.TextContent,
		OrganizationID: body.OrganizationID,
		Active:         body.Active,
		System:         body.System,
		Locale:         body.Locale,
	}
	{
		var zero bool
		if v.Active == zero {
			v.Active = true
		}
	}
	{
		var zero bool
		if v.System == zero {
			v.System = false
		}
	}
	{
		var zero string
		if v.Locale == zero {
			v.Locale = "en"
		}
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]any, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildGetTemplatePayload builds the payload for the email get_template
// endpoint from CLI flags.
func BuildGetTemplatePayload(emailGetTemplateID string, emailGetTemplateJWT string) (*email.GetTemplatePayload, error) {
	var id string
	{
		id = emailGetTemplateID
	}
	var jwt *string
	{
		if emailGetTemplateJWT != "" {
			jwt = &emailGetTemplateJWT
		}
	}
	v := &email.GetTemplatePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildGetTemplateByTypePayload builds the payload for the email
// get_template_by_type endpoint from CLI flags.
func BuildGetTemplateByTypePayload(emailGetTemplateByTypeType string, emailGetTemplateByTypeOrganizationID string, emailGetTemplateByTypeLocale string, emailGetTemplateByTypeJWT string) (*email.GetTemplateByTypePayload, error) {
	var type_ string
	{
		type_ = emailGetTemplateByTypeType
	}
	var organizationID *string
	{
		if emailGetTemplateByTypeOrganizationID != "" {
			organizationID = &emailGetTemplateByTypeOrganizationID
		}
	}
	var locale string
	{
		if emailGetTemplateByTypeLocale != "" {
			locale = emailGetTemplateByTypeLocale
		}
	}
	var jwt *string
	{
		if emailGetTemplateByTypeJWT != "" {
			jwt = &emailGetTemplateByTypeJWT
		}
	}
	v := &email.GetTemplateByTypePayload{}
	v.Type = type_
	v.OrganizationID = organizationID
	v.Locale = locale
	v.JWT = jwt

	return v, nil
}

// BuildUpdateTemplatePayload builds the payload for the email update_template
// endpoint from CLI flags.
func BuildUpdateTemplatePayload(emailUpdateTemplateBody string, emailUpdateTemplateID string, emailUpdateTemplateJWT string) (*email.UpdateTemplatePayload, error) {
	var err error
	var body UpdateTemplateRequestBody
	{
		err = json.Unmarshal([]byte(emailUpdateTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"template\": {\n         \"active\": true,\n         \"html_content\": \"Placeat temporibus molestiae.\",\n         \"locale\": \"Cumque labore delectus dolores.\",\n         \"metadata\": {\n            \"Laudantium magnam corporis iusto.\": \"Et numquam quas praesentium officiis nihil.\"\n         },\n         \"name\": \"Aliquam at aut perferendis et.\",\n         \"subject\": \"Nisi quo.\",\n         \"text_content\": \"Enim mollitia labore repellendus.\"\n      }\n   }'")
		}
		if body.Template == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("template", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var id string
	{
		id = emailUpdateTemplateID
	}
	var jwt *string
	{
		if emailUpdateTemplateJWT != "" {
			jwt = &emailUpdateTemplateJWT
		}
	}
	v := &email.UpdateTemplatePayload{}
	if body.Template != nil {
		v.Template = marshalUpdateEmailTemplateRequestRequestBodyToEmailUpdateEmailTemplateRequest(body.Template)
	}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildDeleteTemplatePayload builds the payload for the email delete_template
// endpoint from CLI flags.
func BuildDeleteTemplatePayload(emailDeleteTemplateID string, emailDeleteTemplateJWT string) (*email.DeleteTemplatePayload, error) {
	var id string
	{
		id = emailDeleteTemplateID
	}
	var jwt *string
	{
		if emailDeleteTemplateJWT != "" {
			jwt = &emailDeleteTemplateJWT
		}
	}
	v := &email.DeleteTemplatePayload{}
	v.ID = id
	v.JWT = jwt

	return v, nil
}

// BuildSendPayload builds the payload for the email send endpoint from CLI
// flags.
func BuildSendPayload(emailSendBody string, emailSendJWT string) (*email.SendPayload, error) {
	var err error
	var body SendRequestBody
	{
		err = json.Unmarshal([]byte(emailSendBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"bcc\": [\n         \"Sunt temporibus similique omnis aperiam.\",\n         \"Voluptatem consequatur facilis pariatur aperiam adipisci error.\",\n         \"Sed consequatur est adipisci veritatis.\",\n         \"Eos itaque.\"\n      ],\n      \"cc\": [\n         \"Mollitia dignissimos ab qui.\",\n         \"Doloribus nulla ut qui sit qui.\",\n         \"Unde unde aspernatur.\"\n      ],\n      \"from\": \"no-reply@example.com\",\n      \"headers\": {\n         \"Qui facere.\": \"Odio quo iure quia ut.\"\n      },\n      \"html_content\": \"Porro aut est adipisci in in consequuntur.\",\n      \"metadata\": {\n         \"Deleniti est ipsum molestiae.\": \"Ut dolore illum.\",\n         \"Rem aut dolores delectus praesentium nobis consectetur.\": \"Est sint ad quam sapiente dolorem reiciendis.\"\n      },\n      \"reply_to\": \"Minus nisi.\",\n      \"subject\": \"Important information\",\n      \"text_content\": \"Architecto culpa dignissimos enim reprehenderit perspiciatis.\",\n      \"to\": [\n         \"user@example.com\"\n      ]\n   }'")
		}
		if body.To == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("to", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if emailSendJWT != "" {
			jwt = &emailSendJWT
		}
	}
	v := &email.SendPayload{
		From:        body.From,
		Subject:     body.Subject,
		HTMLContent: body.HTMLContent,
		TextContent: body.TextContent,
		ReplyTo:     body.ReplyTo,
	}
	if body.To != nil {
		v.To = make([]string, len(body.To))
		for i, val := range body.To {
			v.To[i] = val
		}
	} else {
		v.To = []string{}
	}
	if body.Cc != nil {
		v.Cc = make([]string, len(body.Cc))
		for i, val := range body.Cc {
			v.Cc[i] = val
		}
	}
	if body.Bcc != nil {
		v.Bcc = make([]string, len(body.Bcc))
		for i, val := range body.Bcc {
			v.Bcc[i] = val
		}
	}
	if body.Headers != nil {
		v.Headers = make(map[string]string, len(body.Headers))
		for key, val := range body.Headers {
			tk := key
			tv := val
			v.Headers[tk] = tv
		}
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]any, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}
	v.JWT = jwt

	return v, nil
}

// BuildSendTemplatePayload builds the payload for the email send_template
// endpoint from CLI flags.
func BuildSendTemplatePayload(emailSendTemplateBody string, emailSendTemplateJWT string) (*email.SendTemplatePayload, error) {
	var err error
	var body SendTemplateRequestBody
	{
		err = json.Unmarshal([]byte(emailSendTemplateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"bcc\": [\n         \"Est maiores quis nihil harum.\",\n         \"Illo in similique.\"\n      ],\n      \"cc\": [\n         \"Velit ut.\",\n         \"Ipsa eos neque tempore magnam odio.\",\n         \"Est sit.\",\n         \"Est possimus animi vel.\"\n      ],\n      \"from\": \"no-reply@example.com\",\n      \"headers\": {\n         \"Quis ut sunt.\": \"Non debitis veniam nihil.\"\n      },\n      \"locale\": \"Provident libero est provident consequatur est aliquid.\",\n      \"metadata\": {\n         \"Excepturi assumenda eos fugit quisquam voluptatem.\": \"Blanditiis et placeat quis minus voluptates suscipit.\",\n         \"Rerum animi veritatis provident dolore.\": \"Laborum in aperiam.\"\n      },\n      \"organization_id\": \"Occaecati qui qui.\",\n      \"reply_to\": \"Voluptate voluptate delectus fuga esse deserunt est.\",\n      \"subject\": \"Iste similique voluptatem magnam iste.\",\n      \"template_data\": {\n         \"name\": \"John Doe\"\n      },\n      \"template_type\": \"welcome\",\n      \"to\": [\n         \"user@example.com\"\n      ]\n   }'")
		}
		if body.To == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("to", "body"))
		}
		if body.TemplateData == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("template_data", "body"))
		}
		if err != nil {
			return nil, err
		}
	}
	var jwt *string
	{
		if emailSendTemplateJWT != "" {
			jwt = &emailSendTemplateJWT
		}
	}
	v := &email.SendTemplatePayload{
		From:           body.From,
		Subject:        body.Subject,
		TemplateType:   body.TemplateType,
		OrganizationID: body.OrganizationID,
		Locale:         body.Locale,
		ReplyTo:        body.ReplyTo,
	}
	if body.To != nil {
		v.To = make([]string, len(body.To))
		for i, val := range body.To {
			v.To[i] = val
		}
	} else {
		v.To = []string{}
	}
	if body.TemplateData != nil {
		v.TemplateData = make(map[string]any, len(body.TemplateData))
		for key, val := range body.TemplateData {
			tk := key
			tv := val
			v.TemplateData[tk] = tv
		}
	}
	{
		var zero string
		if v.Locale == zero {
			v.Locale = "en"
		}
	}
	if body.Cc != nil {
		v.Cc = make([]string, len(body.Cc))
		for i, val := range body.Cc {
			v.Cc[i] = val
		}
	}
	if body.Bcc != nil {
		v.Bcc = make([]string, len(body.Bcc))
		for i, val := range body.Bcc {
			v.Bcc[i] = val
		}
	}
	if body.Headers != nil {
		v.Headers = make(map[string]string, len(body.Headers))
		for key, val := range body.Headers {
			tk := key
			tv := val
			v.Headers[tk] = tv
		}
	}
	if body.Metadata != nil {
		v.Metadata = make(map[string]any, len(body.Metadata))
		for key, val := range body.Metadata {
			tk := key
			tv := val
			v.Metadata[tk] = tv
		}
	}
	v.JWT = jwt

	return v, nil
}
