package report

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/threagile/threagile/pkg/types"
)

func WriteRisksJSON(parsedModel *types.Model, filename string) error {
	// Create a structured output that includes category information
	risksWithCategories := make([]map[string]interface{}, 0)

	for _, category := range parsedModel.SortedRiskCategories() {
		risks := parsedModel.SortedRisksOfCategory(category)
		for _, risk := range risks {
			riskData := map[string]interface{}{
				// Original Risk struct fields - keep exact field names
				"category":                         risk.CategoryId,
				"synthetic_id":                     risk.SyntheticId,
				"title":                            risk.Title,
				"severity":                         risk.Severity.String(),
				"exploitation_likelihood":          risk.ExploitationLikelihood.String(),
				"exploitation_impact":              risk.ExploitationImpact.String(),
				"data_breach_probability":          risk.DataBreachProbability.String(),
				"most_relevant_data_asset":         risk.MostRelevantDataAssetId,
				"most_relevant_technical_asset":    risk.MostRelevantTechnicalAssetId,
				"most_relevant_communication_link": risk.MostRelevantCommunicationLinkId,
				"most_relevant_trust_boundary":     risk.MostRelevantTrustBoundaryId,
				"most_relevant_shared_runtime":     risk.MostRelevantSharedRuntimeId,
				"data_breach_technical_assets":     risk.DataBreachTechnicalAssetIDs,
				"risk_status": func() string {
					riskTracking := parsedModel.GetRiskTrackingWithDefault(risk)
					return riskTracking.Status.String()
				}(),

				// Add STRIDE and other category information as new fields
				"stride":               category.STRIDE.String(),
				"stride_title":         category.STRIDE.Title(),
				"category_title":       category.Title,
				"category_description": category.Description,
				"function":             category.Function.String(),
				"function_title":       category.Function.Title(),
				"cwe":                  category.CWE,
				"action":               category.Action,
				"mitigation":           category.Mitigation,
				"check":                category.Check,
			}

			risksWithCategories = append(risksWithCategories, riskData)
		}
	}

	jsonBytes, err := json.Marshal(risksWithCategories)
	if err != nil {
		return fmt.Errorf("failed to marshal risks to JSON: %w", err)
	}
	err = os.WriteFile(filename, jsonBytes, 0600)
	if err != nil {
		return fmt.Errorf("failed to write risks to JSON file: %w", err)
	}
	return nil
}

// TODO: also a "data assets" json?

func WriteTechnicalAssetsJSON(parsedModel *types.Model, filename string) error {
	jsonBytes, err := json.Marshal(parsedModel.TechnicalAssets)
	if err != nil {
		return fmt.Errorf("failed to marshal technical assets to JSON: %w", err)
	}
	err = os.WriteFile(filename, jsonBytes, 0600)
	if err != nil {
		return fmt.Errorf("failed to write technical assets to JSON file: %w", err)
	}
	return nil
}

func WriteStatsJSON(parsedModel *types.Model, filename string) error {
	jsonBytes, err := json.Marshal(overallRiskStatistics(parsedModel))
	if err != nil {
		return fmt.Errorf("failed to marshal stats to JSON: %w", err)
	}
	err = os.WriteFile(filename, jsonBytes, 0600)
	if err != nil {
		return fmt.Errorf("failed to write stats to JSON file: %w", err)
	}
	return nil
}

func overallRiskStatistics(parsedModel *types.Model) riskStatistics {
	result := riskStatistics{}
	result.Risks = make(map[string]map[string]int)
	result.Risks[types.CriticalSeverity.String()] = make(map[string]int)
	result.Risks[types.CriticalSeverity.String()][types.Unchecked.String()] = 0
	result.Risks[types.CriticalSeverity.String()][types.InDiscussion.String()] = 0
	result.Risks[types.CriticalSeverity.String()][types.Accepted.String()] = 0
	result.Risks[types.CriticalSeverity.String()][types.InProgress.String()] = 0
	result.Risks[types.CriticalSeverity.String()][types.Mitigated.String()] = 0
	result.Risks[types.CriticalSeverity.String()][types.FalsePositive.String()] = 0
	result.Risks[types.HighSeverity.String()] = make(map[string]int)
	result.Risks[types.HighSeverity.String()][types.Unchecked.String()] = 0
	result.Risks[types.HighSeverity.String()][types.InDiscussion.String()] = 0
	result.Risks[types.HighSeverity.String()][types.Accepted.String()] = 0
	result.Risks[types.HighSeverity.String()][types.InProgress.String()] = 0
	result.Risks[types.HighSeverity.String()][types.Mitigated.String()] = 0
	result.Risks[types.HighSeverity.String()][types.FalsePositive.String()] = 0
	result.Risks[types.ElevatedSeverity.String()] = make(map[string]int)
	result.Risks[types.ElevatedSeverity.String()][types.Unchecked.String()] = 0
	result.Risks[types.ElevatedSeverity.String()][types.InDiscussion.String()] = 0
	result.Risks[types.ElevatedSeverity.String()][types.Accepted.String()] = 0
	result.Risks[types.ElevatedSeverity.String()][types.InProgress.String()] = 0
	result.Risks[types.ElevatedSeverity.String()][types.Mitigated.String()] = 0
	result.Risks[types.ElevatedSeverity.String()][types.FalsePositive.String()] = 0
	result.Risks[types.MediumSeverity.String()] = make(map[string]int)
	result.Risks[types.MediumSeverity.String()][types.Unchecked.String()] = 0
	result.Risks[types.MediumSeverity.String()][types.InDiscussion.String()] = 0
	result.Risks[types.MediumSeverity.String()][types.Accepted.String()] = 0
	result.Risks[types.MediumSeverity.String()][types.InProgress.String()] = 0
	result.Risks[types.MediumSeverity.String()][types.Mitigated.String()] = 0
	result.Risks[types.MediumSeverity.String()][types.FalsePositive.String()] = 0
	result.Risks[types.LowSeverity.String()] = make(map[string]int)
	result.Risks[types.LowSeverity.String()][types.Unchecked.String()] = 0
	result.Risks[types.LowSeverity.String()][types.InDiscussion.String()] = 0
	result.Risks[types.LowSeverity.String()][types.Accepted.String()] = 0
	result.Risks[types.LowSeverity.String()][types.InProgress.String()] = 0
	result.Risks[types.LowSeverity.String()][types.Mitigated.String()] = 0
	result.Risks[types.LowSeverity.String()][types.FalsePositive.String()] = 0
	for _, risks := range parsedModel.GeneratedRisksByCategory {
		for _, risk := range risks {
			result.Risks[risk.Severity.String()][risk.RiskStatus.String()]++
		}
	}
	return result
}

type riskStatistics struct {
	// TODO add also some more like before / after (i.e. with mitigation applied)
	Risks map[string]map[string]int `yaml:"risks" json:"risks"`
}
