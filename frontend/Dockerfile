FROM node:16.17-alpine3.15

RUN npm install react-scripts@5.0.1 -g --quiet

USER node

# set working directory
WORKDIR /var/www/frontend/

# add `/app/node_modules/.bin` to $PATH
ENV PATH /var/www/frontend/node_modules/.bin:$PATH

# install app dependencies
COPY --chown=node:node package.json ./
COPY --chown=node:node package-lock.json ./
RUN npm install --quiet

RUN chown -R node:node /var/www/frontend/
COPY --chown=node:node . ./

# start app
CMD ["npm", "start", "--env-file", "/var/www/frontend/.env"]

EXPOSE 3000
