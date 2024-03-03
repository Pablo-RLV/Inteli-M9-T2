from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
import time
import itertools

class Chamada():
    def __init__(self):
        self.driver = webdriver.Chrome()
        self.wait = WebDriverWait(self.driver, 10)
        self.client_id = "clientId-"

    def open_page(self):
        self.driver.get("https://www.hivemq.com/demos/websocket-client/")

    def send_first_time(self):
        self.wait.until(EC.presence_of_element_located((By.ID, "clientIdInput"))).clear()
        self.wait.until(EC.presence_of_element_located((By.ID, "clientIdInput"))).send_keys(self.client_id)
        self.wait.until(EC.element_to_be_clickable((By.ID, "connectButton"))).click()

    def loop(self):
        digitos = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
        combinations = itertools.product(digitos, repeat=10)
        for combination in combinations:
            self.client_id = ''.join(combination)
            self.client_id = "clientId-" + self.client_id
            self.publish()
            self.send_consecutive_times()

    def dropdown(self):
        time.sleep(3)
        self.driver.find_element(By.XPATH, "/html/body/div[2]/div[2]/div[3]").click()
    
    def send_consecutive_times(self):
        self.dropdown()
        self.wait.until(EC.element_to_be_clickable((By.ID, "disconnectButton"))).click()
        self.wait.until(EC.presence_of_element_located((By.ID, "clientIdInput"))).clear()
        self.wait.until(EC.presence_of_element_located((By.ID, "clientIdInput"))).send_keys(self.client_id)
        self.wait.until(EC.element_to_be_clickable((By.ID, "connectButton"))).click()
        time.sleep(3)

    def publish(self):
        time.sleep(3)
        self.wait.until(EC.presence_of_element_located((By.ID, "publishTopic"))).clear()
        self.wait.until(EC.presence_of_element_located((By.ID, "publishTopic"))).send_keys("pablin")
        self.wait.until(EC.presence_of_element_located((By.ID, "publishPayload"))).clear()
        self.wait.until(EC.presence_of_element_located((By.ID, "publishPayload"))).send_keys(self.client_id)
        self.wait.until(EC.element_to_be_clickable((By.ID, "publishButton"))).click()

if __name__ == "__main__":
    chamada = Chamada()
    chamada.open_page()
    chamada.send_first_time()
    chamada.loop()