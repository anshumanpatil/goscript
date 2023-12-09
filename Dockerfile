FROM ubuntu:latest
RUN apt-get update && apt-get -y install cron
COPY dailycron /etc/cron.d/dailycron
RUN chmod 0644 /etc/cron.d/dailycron
RUN crontab /etc/cron.d/dailycron
RUN touch /var/log/cron.log
CMD cron && tail -f /var/log/cron.log

