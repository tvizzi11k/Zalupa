FROM python:3.12.3-bookworm
RUN mkdir /code
WORKDIR /code
COPY requirements.txt /code/
RUN apt install gcc pkg-config default-libmysqlclient-dev
RUN pip install -r requirements.txt
COPY . /code/
CMD [ "python", "./megahappy/manage.py", "runserver", "0.0.0.0:8000" ]
