# Step-by-Step Plan

1. **Domain Layer**  
   - Create all entities with validation functions

2. **Repository Interfaces**  
   - Define interfaces for all entities

3. **Usecase Layer**  
   - Implement business logic

4. **Infra Layer**  
   - Implement Postgres repositories  
   - Setup golang-migrate for table creation

5. **Handlers**  
   - Implement HTTP endpoints

6. **Main & Wiring**  
   - Connect everything

7. **Testing**  
   - Unit tests for usecases with mock repos

8. **Optional Enhancements**  
   - Pagination, error handling, logging