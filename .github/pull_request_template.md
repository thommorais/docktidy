## Ticket

<!-- Link to issue/ticket (e.g., #123 or JIRA-456) -->

Closes #

## Summary

<!-- Brief description of what this PR does -->

## How to Review

<!-- Clear instructions for reviewers to see/test the changes -->

### Running the Change

```bash
# Example:
make build
./tmp/docktidy

# Or for specific tests:
make test
go test -v -run TestName ./internal/...
```

### Expected Behavior

<!-- What should the reviewer see/observe? -->

-
-

### Screenshots/Output

<!-- If visual changes or terminal output, include screenshots or paste output here -->
<!-- Remove this section if not applicable -->

```
# Example terminal output
$ ./tmp/docktidy version
docktidy v0.2.0
```

## Checklist

- [ ] Tests added/updated
- [ ] Documentation updated (if needed)
- [ ] Follows conventional commit format
- [ ] No breaking changes (or documented if present)
- [ ] Passes `make check` locally
