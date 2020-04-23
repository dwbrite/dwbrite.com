
.PHONY: host

host:
	-killall dwbrite.com
	sassc root/resources/css/styles.scss root/resources/css/styles.css
	go build -o dwbrite.com src/*
	nohup ./dwbrite.com &

run:
	-killall dwbrite.com
	sassc root/resources/css/styles.scss root/resources/css/styles.css
	go build -o dwbrite.com src/*
	./dwbrite.com
