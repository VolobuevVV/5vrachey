<template>
  <div class="location-map">
    <h2 class="section-title">Как нас найти</h2>

    <div class="location-content">
      <div class="map-wrapper">
        <div class="map-container">
          <iframe
            src="https://yandex.ru/map-widget/v1/?um=constructor%3A28a48935d98f8cda4f9b7a43d3ba6e25e0d55f4bb9ba2a428f6fdc1ff3487d70&amp;lang=ru_RU"
            width="100%"
            height="100%"
            frameborder="0"
          ></iframe>
        </div>
      </div>

      <div class="info-wrapper">
        <div class="info-container">
          <div class="info-content">
            <div class="clinic-name">Пять Врачей в Адлере</div>
            <div class="address">г. Сочи, Адлер, ул. Кирова, 117</div>
            <div class="phone">+7 (862) 123-45-67</div>

            <div class="hours-section">
              <h3 class="hours-title">Часы работы</h3>
              <div class="hours">
                <div class="hours-item">
                  <span class="day">Пн-Пт:</span>
                  <span class="time">08:00 – 20:00</span>
                </div>
                <div class="hours-item">
                  <span class="day">Суббота:</span>
                  <span class="time">09:00 – 18:00</span>
                </div>
                <div class="hours-item">
                  <span class="day">Воскресенье:</span>
                  <span class="time">09:00 – 16:00</span>
                </div>
              </div>
            </div>
          </div>

          <div class="disclaimer">
            <strong>Условия использования</strong>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LocationMap'
}
</script>

<style scoped>
.location-map {
  width: 100%;
  max-width: 1400px;
  margin: 0 auto;
  padding: 2rem;
}

.section-title {
  color: #003449;
  font-size: 1.8rem;
  margin-bottom: 2rem;
  font-weight: 700;
  text-align: center;
}

.location-content {
  display: flex;
  gap: 2rem;
  align-items: stretch;
}

.map-wrapper {
  flex: 1;
}

.map-container {
  width: 100%;
  height: 500px;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 25px rgba(0, 0, 0, 0.08);
}

.map-container iframe {
  border: none;
  display: block;
  width: 100%;
  height: 100%;
}

.info-wrapper {
  width: 400px;
  flex-shrink: 0;
}

.info-container {
  width: 100%;
  height: 100%;
  padding: 2rem;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 2px 25px rgba(0, 0, 0, 0.08);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.info-content {
  flex: 1;
}

.clinic-name {
  font-size: 1.5rem;
  color: #003449;
  margin-bottom: 1rem;
  font-weight: 700;
}

.address {
  font-size: 1rem;
  color: #555;
  margin-bottom: 2rem;
  line-height: 1.5;
}

.phone {
  font-size: 1.4rem;
  color: #003449;
  margin-bottom: 2rem;
  font-weight: 700;
}

.hours-section {
  margin-top: 2rem;
}

.hours-title {
  color: #003449;
  margin-bottom: 1rem;
  font-size: 1.1rem;
  border-bottom: 2px solid #e0e0e0;
  padding-bottom: 0.5rem;
  font-weight: 600;
}

.hours {
  font-size: 1rem;
}

.hours-item {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.8rem;
  padding: 0.5rem 0;
  border-bottom: 1px solid #f1f3f4;
}

.day {
  color: #555;
  font-weight: 500;
}

.time {
  color: #003449;
  font-weight: 600;
}

.disclaimer {
  margin-top: 2rem;
  font-size: 0.8rem;
  color: #777;
  border-top: 1px solid #e0e0e0;
  padding-top: 1rem;
  text-align: center;
}
</style>
