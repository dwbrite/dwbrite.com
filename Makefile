
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

service: 
	sudo cp -f dwbrite_com.service /etc/systemd/system/dwbrite_com.service
	sudo systemctl enable dwbrite_com.service
	sudo systemctl restart dwbrite_com.service

# todo: nginx
