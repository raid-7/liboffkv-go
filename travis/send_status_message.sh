#!/bin/bash

if [[ -z "$TELEGRAM_TOKEN" ]] || [[ -z "$TELEGRAM_CHAT" ]]; then
	exit 1
fi

cd "$HOME"

if [[ "$TRAVIS_TEST_RESULT" == "0" ]]; then
	# success
MESSAGE="Go implementation

*Success:* ${TRAVIS_JOB_WEB_URL}
*Commit*: _${TRAVIS_COMMIT_MESSAGE}_ into \`${TRAVIS_BRANCH}\`
*System*: \`${TRAVIS_OS_NAME}\`"

else
	#failure
MESSAGE="Go implementation

*Failure:* ${TRAVIS_JOB_WEB_URL}
*Commit*: _${TRAVIS_COMMIT_MESSAGE}_ into \`${TRAVIS_BRANCH}\`
*System*: \`${TRAVIS_OS_NAME}\`"

fi

curl -X POST \
	-F "chat_id=${TELEGRAM_CHAT}" -F "parse_mode=Markdown" -F "disable_web_page_preview=true" -F \
	"text=${MESSAGE}" https://api.telegram.org/bot${TELEGRAM_TOKEN}/sendMessage
