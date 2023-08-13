FROM nginx:latest

EXPOSE 80

RUN apt-get update -qq && apt-get install apt-transport-https curl vim procps -y -qq

COPY conf.d /etc/nginx/conf.d
COPY nginx.conf /etc/nginx/nginx.conf

# ENTRYPOINT ["/docker-entrypoint.sh"]
