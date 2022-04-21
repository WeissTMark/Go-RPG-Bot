# pull official base image
FROM node:16.14.2-alpine
# set working directory
WORKDIR /app
ENV CHOKIDAR_USEPOLLING = true

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

# add app
COPY . ./

# start app
CMD ["npm", "start"]