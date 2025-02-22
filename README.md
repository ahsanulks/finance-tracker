# finance-tracker

# Solution
## Requirement
### 1. Functional Requirements
The application should be able to:
1. Read transaction data from a user-specified CSV file.
2. Filter transactions by the specified period (`YYYYMM`). Only transactions within the given year and month should be processed.
3. Calculate and return:
   - Period should follow the format `YYYYMM`
   - Total income
   - Total expenditure
   - List of transactions, including:
     - Date
     - Amount
     - Content
     - Transactions should be sorted in descending order by date.
4. Output the results in JSON format to stdout.
5. Expect an input CSV file with the following format:
   - The date is string should follow the format `YYYY/MM/DD`.
   - The date of occurrence of deposits and withdrawals is not limited to a specific year and month
   - The amount is a integer that describes amount of deposit/withdrawal.
   - The content is a string that describes the transaction.
   - None of the columns are allowed to be empty.
   - The order of rows is not guaranteed.
```csv
date,amount,content
2023/06/15,-720,transportation
2022/01/05,-1000,eating out
2022/01/06,-10000,debit
2022/02/03,-1500,dining out
2022/01/25,-100000,rent
2023/03/01,200000,salary
```

### 2. Non-Functional Requirements
1. The program should be designed with flexibility in mind to support future enhancements, such as exporting data to a file or supporting different output formats.
3. The application must be efficient and capable of handling large CSV files without consuming excessive memory, ensuring it does not impact system performance.
4. It should provide clear error handling and validation for incorrect inputs, such as:
   - Invalid date format for the period.
   - Missing or unreadable CSV files.
5. The program should have a simple and intuitive command-line interface.

### 3. Additional Considerations
- Handling large CSV files: since there is no file size limit, the program should process the CSV data efficiently using streaming or line-by-line reading instead of loading the entire file into memory.
- Data format assumptions:
  - The CSV file follows a structured format.
  - The amount field uses negative values for expenses and positive values for income.

## Workflow
```mermaid
flowchart TB
    A[Understand the Problem] --> B[Analyze Constraints & Conditions]
    B --> C[Define Functional/Non-Functional Requirements and Additional Considerations]
```
1. Understand the Problem: Carefully read and analyze the problem statement. Ensure all key details are noted.
2. Analyze Constraints & Conditions: Identify any specific constraints, assumptions, or conditions that the solution must adhere to.
3. Define Functional/Non-Functional Requirements and Additional Considerations: List the essential features and behaviors the application must support. Look for edge cases, ambiguities, or undefined scenarios that may impact the solution.
