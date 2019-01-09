FROM alpine:latest
LABEL maintainer="jeffspahr"
RUN mkdir /opt/logger \
  && chgrp -R 0 /opt/logger \
  && chmod -R g+rwX /opt/logger
WORKDIR /opt/logger
ADD . /opt/logger
CMD ["/opt/logger/logger"]
