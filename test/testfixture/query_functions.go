package testfixture

import (
	"github.com/fabric8-services/fabric8-wit/account"
	"github.com/fabric8-services/fabric8-wit/area"
	"github.com/fabric8-services/fabric8-wit/iteration"
	"github.com/fabric8-services/fabric8-wit/label"
	"github.com/fabric8-services/fabric8-wit/query"
	"github.com/fabric8-services/fabric8-wit/spacetemplate"
	"github.com/fabric8-services/fabric8-wit/workitem"
	"github.com/fabric8-services/fabric8-wit/workitem/link"
	errs "github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// LabelByName returns the first label that has the given name (if any). If you
// have labels with the same name in different spaces you can also pass in one
// space ID to filter by space as well.
func (fxt *TestFixture) LabelByName(name string, spaceID ...uuid.UUID) *label.Label {
	for _, l := range fxt.Labels {
		if l.Name == name && len(spaceID) > 0 && l.SpaceID == spaceID[0] {
			return l
		} else if l.Name == name && len(spaceID) == 0 {
			return l
		}
	}
	return nil
}

// IterationByName returns the first iteration that has the given name (if any).
// If you have iterations with the same name in different spaces you can also
// pass in one space ID to filter by space as well.
func (fxt *TestFixture) IterationByName(name string, spaceID ...uuid.UUID) *iteration.Iteration {
	for _, i := range fxt.Iterations {
		if i.Name == name && len(spaceID) > 0 && i.SpaceID == spaceID[0] {
			return i
		} else if i.Name == name && len(spaceID) == 0 {
			return i
		}
	}
	return nil
}

// AreaByName returns the first area that has the given name (if any). If you
// have areas with the same name in different spaces you can also pass in one
// space ID to filter by space as well.
func (fxt *TestFixture) AreaByName(name string, spaceID ...uuid.UUID) *area.Area {
	for _, a := range fxt.Areas {
		if a.Name == name && len(spaceID) > 0 && a.SpaceID == spaceID[0] {
			return a
		} else if a.Name == name && len(spaceID) == 0 {
			return a
		}
	}
	return nil
}

// WorkItemTypeByName returns the first work item type that has the given name
// (if any). If you have work item types with the same name in different space
// templates you can also pass in one space template ID to filter by space
// template as well.
func (fxt *TestFixture) WorkItemTypeByName(name string, spaceTemplateID ...uuid.UUID) *workitem.WorkItemType {
	for _, wit := range fxt.WorkItemTypes {
		if wit.Name == name && len(spaceTemplateID) > 0 && wit.SpaceTemplateID == spaceTemplateID[0] {
			return wit
		} else if wit.Name == name && len(spaceTemplateID) == 0 {
			return wit
		}
	}
	return nil
}

// WorkItemTypeByID returns the work item type that has the given ID (if any).
func (fxt *TestFixture) WorkItemTypeByID(id uuid.UUID) *workitem.WorkItemType {
	for _, wit := range fxt.WorkItemTypes {
		if wit.ID == id {
			return wit
		}
	}
	return nil
}

// IdentityByUsername returns the first identity that has the given username (if
// any).
func (fxt *TestFixture) IdentityByUsername(username string) *account.Identity {
	for _, i := range fxt.Identities {
		if i.Username == username {
			return i
		}
	}
	return nil
}

// WorkItemByTitle returns the first work item that has the given title (if
// any). If you have work items with the same title in different spaces you can
// also pass in one space ID to filter by space as well.
func (fxt *TestFixture) WorkItemByTitle(title string, spaceID ...uuid.UUID) *workitem.WorkItem {
	for _, wi := range fxt.WorkItems {
		v, ok := wi.Fields[workitem.SystemTitle]
		if !ok {
			panic(errs.Errorf("failed to find work item with title '%s' (field '%s' does not exist in work item title)", title, workitem.SystemTitle))
		}
		if v == title && len(spaceID) > 0 && wi.SpaceID == spaceID[0] {
			return wi
		} else if v == title && len(spaceID) == 0 {
			return wi
		}
	}
	return nil
}

// WorkItemByID returns the first work item that has the given ID (if any).
func (fxt *TestFixture) WorkItemByID(ID uuid.UUID) *workitem.WorkItem {
	for _, wi := range fxt.WorkItems {
		if ID == wi.ID {
			return wi
		}
	}
	return nil
}

// WorkItemLinkTypeByName returns the first work item link type that has the
// given name (if any). If you have work item link types with the same name in
// different space templates you can also pass in one space template ID to
// filter by space template as well.
func (fxt *TestFixture) WorkItemLinkTypeByName(name string, spaceTemplateID ...uuid.UUID) *link.WorkItemLinkType {
	for _, wilt := range fxt.WorkItemLinkTypes {
		if wilt.Name == name && len(spaceTemplateID) > 0 && wilt.SpaceTemplateID == spaceTemplateID[0] {
			return wilt
		} else if wilt.Name == name && len(spaceTemplateID) == 0 {
			return wilt
		}
	}
	return nil
}

// QueryByTitle returns the first query that has the given title (if
// any). If you have queries with the same title in different spaces you can
// also pass in one space ID to filter by space as well.
func (fxt *TestFixture) QueryByTitle(title string, spaceID ...uuid.UUID) *query.Query {
	for _, q := range fxt.Queries {
		if q.Title == title && len(spaceID) > 0 && q.SpaceID == spaceID[0] {
			return q
		} else if q.Title == title && len(spaceID) == 0 {
			return q
		}
	}
	return nil
}

// SpaceTemplateByName returns the first space template that has the given name (if any).
func (fxt *TestFixture) SpaceTemplateByName(name string) *spacetemplate.SpaceTemplate {
	for _, st := range fxt.SpaceTemplates {
		if st.Name == name {
			return st
		}
	}
	return nil
}
