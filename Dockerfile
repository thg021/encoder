FROM golang:1.14-alpine3.11
ENV PATH="$PATH:/bin/bash" \
  BENTO4_BIN="/opt/bento4/bin" \
  PATH="$PATH:/opt/bento4/bin"

# FFMPEG
RUN apk add --update ffmpeg bash curl make

# Install Bento
WORKDIR /tmp/bento4
ENV BENTO4_BASE_URL="http://zebulon.bok.net/Bento4/source/" \
  BENTO4_VERSION="1-5-0-615" \
  BENTO4_CHECKSUM="5378dbb374343bc274981d6e2ef93bce0851bda1" \
  BENTO4_TARGET="" \
  BENTO4_PATH="/opt/bento4" \
  BENTO4_TYPE="SRC"
# download and unzip bento4
RUN apk add --update --upgrade curl python unzip bash gcc g++ scons && \
  curl -O -s ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
  sha1sum -b Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip | grep -o "^$BENTO4_CHECKSUM " && \
  mkdir -p ${BENTO4_PATH} && \
  unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip -d ${BENTO4_PATH} && \
  rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
  apk del unzip && \
  # don't do these steps if using binary install
  cd ${BENTO4_PATH} && scons -u build_config=Release target=x86_64-unknown-linux && \
  cp -R ${BENTO4_PATH}/Build/Targets/x86_64-unknown-linux/Release ${BENTO4_PATH}/bin && \
  cp -R ${BENTO4_PATH}/Source/Python/utils ${BENTO4_PATH}/utils && \
  cp -a ${BENTO4_PATH}/Source/Python/wrappers/. ${BENTO4_PATH}/bin

# utitilizar o git e zsh 
RUN apk add --update --upgrade \
  git \
  ca-certificates \
  zsh \
  wget \
  procps

RUN git clone https://github.com/powerline/fonts.git --depth=1

RUN ./fonts/install.sh
RUN rm -rf fonts

WORKDIR /go/src

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.2/zsh-in-docker.sh)" -- \
  -t https://github.com/romkatv/powerlevel10k \
  -p git \
  -p git-flow \
  -p https://github.com/zdharma-continuum/fast-syntax-highlighting \
  -p https://github.com/zsh-users/zsh-autosuggestions \
  -p https://github.com/zsh-users/zsh-completions \
  -a 'export TERM=xterm-256color'

RUN echo '[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh' >> ~/.zshrc && \
  echo 'HISTFILE=/go/.zsh_history' >> ~/.zshrc



#vamos mudar para o endpoint correto. Usando top apenas para segurar o processo rodando
CMD [ "tail", "-f" , "/dev/null" ]
