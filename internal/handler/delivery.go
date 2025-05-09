package handler

import (
	"encoding/json"
	"net/http"

	"greedy-game/internal/db/sqlc"
	"greedy-game/internal/models"

)

type Handler struct {
	Queries *sqlc.Queries
}

func NewHandler(q *sqlc.Queries) *Handler {
	return &Handler{Queries: q}
}

func (h *Handler) Delivery(w http.ResponseWriter, r *http.Request) {
	app := r.URL.Query().Get("app")
	country := r.URL.Query().Get("country")
	os := r.URL.Query().Get("os")

	if app == "" || country == "" || os == "" {
		http.Error(w, `{"error":"missing app/country/os param"}`, http.StatusBadRequest)
		return
	}

	campaigns, err := h.Queries.GetActiveCampaigns(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rules, err := h.Queries.GetTargetingRules(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	matching := []model.Campaign{}
	for _, rule := range rules {
		if matchesRule(rule, app, os, country) {
			for _, camp := range campaigns {
				if camp.ID == rule.CampaignID.String {
					matching = append(matching, model.Campaign{
						ID:    camp.ID,
						Image: camp.Image,
						CTA:   camp.Cta,
					})
				}
			}
		}
	}

	if len(matching) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(matching)
}

func matchesRule(rule sqlc.TargetingRule, app, os, country string) bool {
	if len(rule.IncludeApps) > 0 && !contains(rule.IncludeApps, app) {
		return false
	}
	if contains(rule.ExcludeApps, app) {
		return false
	}
	if len(rule.IncludeOs) > 0 && !contains(rule.IncludeOs, os) {
		return false
	}
	if contains(rule.ExcludeOs, os) {
		return false
	}
	if len(rule.IncludeCountries) > 0 && !contains(rule.IncludeCountries, country) {
		return false
	}
	if contains(rule.ExcludeCountries, country) {
		return false
	}
	return true
}

func contains(arr []string, val string) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}
