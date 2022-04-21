FROM node:14-alpine3.12 as base

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
ENV PATH /usr/src/app/node_modules/.bin:$PATH
ENV CHOKIDAR_USEPOLLING true

COPY package.json /usr/src/app
RUN yarn install

CMD ["yarn", "start"]
#ENTRYPOINT ["tail"]
#CMD ["-f","/dev/null"]