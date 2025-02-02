// Code generated by enthistory, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package generated

import "github.com/datumforge/enthistory"

// withHistory adds the history hooks to the appropriate schemas - generated by enthistory
func (c *Client) WithHistory() {
	for _, hook := range enthistory.HistoryHooks[*ContactMutation]() {
		c.Contact.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*DocumentDataMutation]() {
		c.DocumentData.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*EntitlementMutation]() {
		c.Entitlement.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*EntitlementPlanMutation]() {
		c.EntitlementPlan.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*EntitlementPlanFeatureMutation]() {
		c.EntitlementPlanFeature.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*EntityMutation]() {
		c.Entity.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*EntityTypeMutation]() {
		c.EntityType.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*EventMutation]() {
		c.Event.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*FeatureMutation]() {
		c.Feature.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*FileMutation]() {
		c.File.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*GroupMutation]() {
		c.Group.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*GroupMembershipMutation]() {
		c.GroupMembership.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*GroupSettingMutation]() {
		c.GroupSetting.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*HushMutation]() {
		c.Hush.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*IntegrationMutation]() {
		c.Integration.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*NoteMutation]() {
		c.Note.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*OauthProviderMutation]() {
		c.OauthProvider.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*OrgMembershipMutation]() {
		c.OrgMembership.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*OrganizationMutation]() {
		c.Organization.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*OrganizationSettingMutation]() {
		c.OrganizationSetting.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*TemplateMutation]() {
		c.Template.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*UserMutation]() {
		c.User.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*UserSettingMutation]() {
		c.UserSetting.Use(hook)
	}
	for _, hook := range enthistory.HistoryHooks[*WebhookMutation]() {
		c.Webhook.Use(hook)
	}
}
