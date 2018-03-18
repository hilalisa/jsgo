package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type ModalBuilder struct {
	id   string
	save func(*vecty.Event)
	body []vecty.MarkupOrChild
}

func Modal(id string, save func(*vecty.Event)) *ModalBuilder {
	return &ModalBuilder{
		id:   id,
		save: save,
	}
}

func (m *ModalBuilder) Body(body ...vecty.MarkupOrChild) *ModalBuilder {
	m.body = body
	return m
}

func (m *ModalBuilder) Build() vecty.ComponentOrHTML {

	body := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.Class("modal-body"),
		),
	}
	body = append(body, m.body...)

	return elem.Div(
		vecty.Markup(
			prop.ID(m.id),
			vecty.Class("modal"),
			vecty.Property("tabindex", "-1"),
			vecty.Property("role", "dialog"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("modal-dialog"),
				vecty.Property("role", "dialog"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("modal-content"),
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("modal-header"),
					),
					elem.Heading5(
						vecty.Markup(
							vecty.Class("modal-title"),
						),
						vecty.Text("Add file"),
					),
					elem.Button(
						vecty.Markup(
							prop.Type(prop.TypeButton),
							vecty.Class("close"),
							vecty.Data("dismiss", "modal"),
							vecty.Property("aria-label", "Close"),
						),
						elem.Span(
							vecty.Markup(
								vecty.Property("aria-hidden", "true"),
							),
							vecty.Text("×"),
						),
					),
				),
				elem.Div(
					body...,
				),
				elem.Div(
					vecty.Markup(
						vecty.Class("modal-footer"),
					),
					elem.Button(
						vecty.Markup(
							prop.Type(prop.TypeButton),
							vecty.Class("btn", "btn-primary"),
							event.Click(m.save).PreventDefault(),
						),
						vecty.Text("Save"),
					),
					elem.Button(
						vecty.Markup(
							prop.Type(prop.TypeButton),
							vecty.Class("btn", "btn-secondary"),
							vecty.Data("dismiss", "modal"),
						),
						vecty.Text("Close"),
					),
				),
			),
		),
	)
}
