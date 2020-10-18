all:

TAG = 1.0
PREFIX = briandtb/nginxhttps
KEY = nginx.key
CERT = nginx.crt
SECRET = nginxsecret.yaml

keys:
	# The CName used here is specific to the service specified in nginx-app.yaml
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout $(KEY) -out $(CERT) -subj "/CN=nginxsvc/O=nginxsvc"

container:
	docker build --pull -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean: 
	rm $(KEY)
	rm $(CERT)

secret:
	./create-secret.sh $(KEY) $(CERT) $(SECRET) 
