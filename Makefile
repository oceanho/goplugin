.PHONY: plugins

.PHONY: auth oauth2 clean

plugins: ldap oauth2

ldap:
	@go build --buildmode=plugin -o build/plugins/auth/ldap.so plugins/impl/auth/ldap/ldap.go

oauth2:
	@go build --buildmode=plugin -o build/plugins/auth/oauth2.so plugins/impl/auth/oauth2/oauth2.go

clean:
	@rm -rf build/
