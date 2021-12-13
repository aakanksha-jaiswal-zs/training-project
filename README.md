# training-project

Quick set up of the project structure:
```
mkdir sample-api && cd sample-api 
go mod init sample-api
touch main.go
mkdir -p handlers/student services/student stores/student
mkdir -p handlers/employee services/employee stores/employee
cd services && touch interface.go && cd ..
cd stores && touch interface.go && cd ..
mkdir models && cd models && touch student.go employee.go && cd ..
mkdir filters && cd models && touch student.go employee.go && cd ..
touch handlers/student/student.go handlers/student/student_test.go
touch handlers/employee/employee.go handlers/employee/employee_test.go
touch services/student/student.go services/student/student_test.go
touch services/employee/employee.go services/employee/employee_test.go
touch stores/employee/employee.go stores/employee/employee_test.go    
touch stores/student/student.go stores/student/student_test.go
mkdir -p errors drivers middlewares
touch .golangci.yml .gitignore
mkdir -p errors drivers middlewares
touch drivers/sql.go middlewares/auth.go middlewares/log.go
cd errors && touch db.go entity_already_exists.go entity_not_found.go invalid_param.go missing_param.go && cd ..
```

Initialize a git repository:
```
git init .
git remote add origin <origin>
```
Branch naming conventions:
```
(Format)
<group-token>/<issue-no>/<description>
```
1. group token- feature FE, enhancement EN, bug fix BUG, documentation DOC, temporary TMP
2. issue number
3. description (small case, separated by hyphens)

for example, we can have a branch named as:
```
FE/1302/add-dynamodb-migartion-support
EN/12/improve-code-quality
BUG/121/fix-race-condition
DOC/102/document-cicd-setup
TMP/1/test-sample-api-workflow
```

Refer [this](https://cbea.ms/git-commit/) for a guide to writing proper commit messages