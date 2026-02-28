# Load environment variables from .env_dev file
ifneq (,$(wildcard ./.env_dev))
    include .env_dev
    export
endif

help:
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)
server:			## Starts backend server with hot reload
	air
sync_assets:		## Watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
	air \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--build.include_ext "js,css"
sync_templ:		## Sync templ generated files
	templ generate --watch --proxy="http://localhost:$(SERVER_PORT)" --open-browser=false
build_css:		## Builds static css
	npx @tailwindcss/cli -i modules/shared/web/css/base.css -o public/styles.css --watch
dev:			## Starts a dev server
	make -j4 server sync_templ build_css sync_assets
