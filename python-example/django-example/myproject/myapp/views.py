# Create your views here.
from django.shortcuts import render
from django.http import HttpResponse
from datetime import datetime

def hello(request):
	text = "<html><body>hello world!</body></html>"
	return HttpResponse(text)

def morning(request):
	return render(request, "morning.html", {})

def param(request, param):
	text = "<html><body>Param is %s </body></html>"%param
	return HttpResponse(text)

def date_display(request, month, year):
	text = "<html><body>The date is %s/%s</body></html>"%(month, year)
	return HttpResponse(text)

def today(request):
	date = datetime.now().date()
	daysOfWeek = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
	return render(request, "today.html", {"today": date, "daysOfWeek": daysOfWeek})
