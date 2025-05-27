# Threagile TODO List

## Container Implementation
- Set up the container to work right away and allow checking changes within it

## Code Changes
- Implement the mapping from Elevated to High severity **[DONE]**

## Report Changes
- Remove `Elevated` severity from reports

## Data Assets **[DONE]**
- Add data_classification tag with the following possible values:
  - Public
  - Internal Use
  - Confidential
  - Restricted
  - No Data Stored

- Add personal_data tag with the following possible values:
  - Sensitive Personal Data
  - Personal Data (including personally identifiable information PII)
  - Non-Public Information (NPI)
  - None
  - Unknown

## CVSS Integration
- Calculate Vector String and Rating according to [NIST CVSS v3 calculator](https://nvd.nist.gov/vuln-metrics/cvss/v3-calculator)
  - We threagile ustawić te metryki, które się da 
    - Attack Complexity (AC) -> Likelyhood?
    - Privileges Required (PR) -> tam gdzie comlink wymaga autoryzacji to jest high
    - Scope (S) -> Impact
  - W rozszerzeniu Do VsCode dodać możliwość ustawienia wszystkich metryk i tam obliczać według kalkulatora
  

## Technical Asset Enhancements **[DONE]**
- Modify the `technologies` tag by changing it to an array of objects that have the following properties:
  - `name`: "POSTMAN"
  - `version`: 4.0
  - `description`
- Add a `trust_level` tag which allows the following values:
  - "Valid User Credentials"
  - "API Key"
  - "Service Account"
- Add new boolean tag `internally_developed`
- schema do modyfikacji