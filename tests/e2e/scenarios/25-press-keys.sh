#!/bin/bash
# 25-press-keys.sh — Verify press action sends actual key events (not text)
#
# Regression test for GitHub issue #236: press action was typing key names
# as literal text instead of dispatching keyboard events.

source "$(dirname "$0")/common.sh"

# ─────────────────────────────────────────────────────────────────
start_test "press Enter: submits form (not types 'Enter')"

# Navigate to form fixture
pt_post /navigate -d "{\"url\":\"${FIXTURES_URL}/form.html\"}"
sleep 1

# Type into username field
pt_post /action -d '{"kind":"type","selector":"#username","text":"testuser"}'
assert_ok "type into username"

# Press Enter to submit the form
pt_post /action -d '{"kind":"press","key":"Enter"}'
assert_ok "press Enter"

# Give form submit handler time to execute
sleep 0.5

# Check that "Form submitted!" appears (proves Enter triggered submit)
pt_get /snapshot
if echo "$RESULT" | grep -q "Form submitted"; then
  echo -e "  ${GREEN}✓${NC} form was submitted (Enter key worked)"
  ((ASSERTIONS_PASSED++)) || true
else
  echo -e "  ${RED}✗${NC} form was NOT submitted - Enter may have typed as text"
  ((ASSERTIONS_FAILED++)) || true
fi

# Check that username field does NOT contain "Enter" as text
pt_post /evaluate -d '{"expression":"document.getElementById(\"username\").value"}'
USERNAME_VALUE=$(echo "$RESULT" | jq -r '.result // empty')
if echo "$USERNAME_VALUE" | grep -qi "enter"; then
  echo -e "  ${RED}✗${NC} username contains 'Enter' text: $USERNAME_VALUE (bug #236)"
  ((ASSERTIONS_FAILED++)) || true
else
  echo -e "  ${GREEN}✓${NC} username field clean: $USERNAME_VALUE"
  ((ASSERTIONS_PASSED++)) || true
fi

end_test

# ─────────────────────────────────────────────────────────────────
start_test "press Tab: moves focus (not types 'Tab')"

# Navigate fresh
pt_post /navigate -d "{\"url\":\"${FIXTURES_URL}/form.html\"}"
sleep 1

# Focus on username and type
pt_post /action -d '{"kind":"click","selector":"#username"}'
pt_post /action -d '{"kind":"type","selector":"#username","text":"hello"}'
assert_ok "type hello"

# Press Tab to move to next field
pt_post /action -d '{"kind":"press","key":"Tab"}'
assert_ok "press Tab"

# Verify username doesn't contain "Tab" text
pt_post /evaluate -d '{"expression":"document.getElementById(\"username\").value"}'
USERNAME_VALUE=$(echo "$RESULT" | jq -r '.result // empty')
if echo "$USERNAME_VALUE" | grep -qi "tab"; then
  echo -e "  ${RED}✗${NC} username contains 'Tab' text: $USERNAME_VALUE (bug #236)"
  ((ASSERTIONS_FAILED++)) || true
else
  echo -e "  ${GREEN}✓${NC} username field has no 'Tab' text: $USERNAME_VALUE"
  ((ASSERTIONS_PASSED++)) || true
fi

end_test
