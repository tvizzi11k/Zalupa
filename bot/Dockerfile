FROM python:3.12.3-bookworm
RUN mkdir /code
WORKDIR /code
RUN apt install gcc pkg-config default-libmysqlclient-dev
COPY requirements.txt /code/
RUN pip --disable-pip-version-check install --no-compile -r requirements.txt
COPY . /code/
CMD [ "python", "bot.py" ]
