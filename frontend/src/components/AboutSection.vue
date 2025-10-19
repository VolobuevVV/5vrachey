<template>
  <section class="about-section">
    <div class="about-container">

      <!-- Первая ячейка - О медцентре -->
      <div class="about-item about-item-main">
        <div class="about-content">
          <div class="about-text">
            <h2 class="about-title">Медицинский центр "Пять врачей"</h2>
            <p class="about-description">
              Современный многопрофильный медицинский центр, объединяющий лучших специалистов
              и передовые технологии для комплексной заботы о вашем здоровье.
            </p>
            <div class="about-features">
              <div class="feature">
                <span class="feature-number">15+</span>
                <span class="feature-text">лет опыта</span>
              </div>
              <div class="feature">
                <span class="feature-number">5000+</span>
                <span class="feature-text">довольных пациентов</span>
              </div>
              <div class="feature">
                <span class="feature-number">20+</span>
                <span class="feature-text">врачей экспертов</span>
              </div>
            </div>
          </div>
          <div class="about-images">
            <div class="image-main">
              <div class="image-placeholder"></div>
            </div>
            <div class="image-secondary">
              <div class="image-placeholder"></div>
            </div>
          </div>
        </div>
      </div>

      <!-- Вторая ячейка - Оборудование -->
      <div class="about-item about-item-equipment">
        <div class="about-content reverse">
          <div class="about-text">
            <h2 class="about-title">Современное оборудование</h2>
            <p class="about-description">
              Мы используем только современное диагностическое и лечебное оборудование
              от ведущих мировых производителей, что гарантирует точность диагноза
              и эффективность лечения.
            </p>
            <ul class="equipment-list">
              <li>Цифровые системы диагностики</li>
              <li>Ультразвуковые аппараты экспертного класса</li>
              <li>Лабораторные анализаторы последнего поколения</li>
              <li>Эндоскопические системы высокой четкости</li>
            </ul>
          </div>
          <div class="about-images">
            <div class="image-stack">
              <div class="image-top">
                <div class="image-placeholder"></div>
              </div>
              <div class="image-bottom">
                <div class="image-placeholder"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Третья ячейка - Подход -->
      <div class="about-item about-item-approach">
        <div class="about-content">
          <div class="about-text">
            <h2 class="about-title">Индивидуальный подход</h2>
            <p class="about-description">
              Мы считаем, что каждый пациент уникален. Наши врачи разрабатывают
              персонализированные программы лечения и профилактики, учитывая
              все особенности вашего здоровья.
            </p>
            <div class="approach-highlights">
              <div class="highlight">
                <div class="highlight-icon"></div>
                <span>Комплексная диагностика</span>
              </div>
              <div class="highlight">
                <div class="highlight-icon"></div>
                <span>Междисциплинарный консилиум</span>
              </div>
              <div class="highlight">
                <div class="highlight-icon"></div>
                <span>Сопровождение на всех этапах</span>
              </div>
            </div>
          </div>
          <div class="about-images">
            <div class="image-overlap">
              <div class="image-back">
                <div class="image-placeholder"></div>
              </div>
              <div class="image-front">
                <div class="image-placeholder"></div>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>
  </section>
</template>

<script>
export default {
  name: 'AboutSection'
}
</script>

<style scoped>
.about-section {
  margin-top: 60px;
  padding: 4rem 0;
  background: linear-gradient(135deg, #f8fdff 0%, #ffffff 100%);
}

.about-container {
  max-width: 1200px;
  margin: 0 auto;
}

.about-item {
  margin-bottom: 6rem;
}

.about-item:last-child {
  margin-bottom: 0;
}

.about-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: center;
}

.about-content.reverse {
  grid-template-columns: 1fr 1fr;
  direction: rtl;
}

.about-content.reverse .about-text {
  direction: ltr;
}

.about-text {
  padding: 2rem;
}

.about-title {
  color: rgb(0, 52, 73);
  font-size: 2.2rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  line-height: 1.2;
}

.about-description {
  color: #2c3e50;
  font-size: 1.1rem;
  line-height: 1.6;
  margin-bottom: 2rem;
}

/* Первая ячейка - О медцентре */
.about-features {
  display: flex;
  gap: 2rem;
  margin-top: 2rem;
}

.feature {
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.feature-number {
  color: rgb(6, 113, 133);
  font-size: 2rem;
  font-weight: 700;
  line-height: 1;
}

.feature-text {
  color: #003449;
  font-size: 0.9rem;
  font-weight: 600;
  margin-top: 0.5rem;
}

.about-images {
  position: relative;
  height: 400px;
}

.image-main {
  position: absolute;
  top: 0;
  left: 0;
  width: 70%;
  height: 80%;
  z-index: 2;
}

.image-secondary {
  position: absolute;
  bottom: 0;
  right: 0;
  width: 60%;
  height: 70%;
  z-index: 1;
}

/* Вторая ячейка - Оборудование */
.equipment-list {
  list-style: none;
  padding: 0;
  margin: 2rem 0;
}

.equipment-list li {
  color: #2c3e50;
  padding: 0.5rem 0;
  padding-left: 1.5rem;
  position: relative;
  font-size: 1rem;
}

.equipment-list li::before {
  content: '';
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 8px;
  height: 8px;
  background: rgb(6, 113, 133);
  border-radius: 50%;
}

.image-stack {
  position: relative;
  height: 400px;
}

.image-top {
  position: absolute;
  top: 0;
  right: 0;
  width: 70%;
  height: 60%;
  z-index: 2;
  transform: rotate(-2deg);
}

.image-bottom {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 75%;
  height: 65%;
  z-index: 1;
  transform: rotate(1deg);
}

/* Третья ячейка - Подход */
.approach-highlights {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-top: 2rem;
}

.highlight {
  display: flex;
  align-items: center;
  gap: 1rem;
  color: #2c3e50;
  font-weight: 500;
}

.highlight-icon {
  width: 12px;
  height: 12px;
  background: rgb(6, 113, 133);
  border-radius: 50%;
  flex-shrink: 0;
}

.image-overlap {
  position: relative;
  height: 400px;
}

.image-back {
  position: absolute;
  top: 0;
  left: 10%;
  width: 80%;
  height: 70%;
  z-index: 1;
}

.image-front {
  position: absolute;
  bottom: 0;
  right: 5%;
  width: 70%;
  height: 80%;
  z-index: 2;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

/* Заглушки для изображений */
.image-placeholder {
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #e3f2fd 0%, #f3e5f5 100%);
  border: 2px solid rgb(6, 113, 133);
  border-radius: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: rgb(6, 113, 133);
  font-weight: 600;
  font-size: 1rem;
}

.image-main .image-placeholder {
  background: linear-gradient(135deg, #e3f2fd 0%, #bbdefb 100%);
}

.image-secondary .image-placeholder {
  background: linear-gradient(135deg, #f3e5f5 0%, #e1bee7 100%);
}

.image-top .image-placeholder {
  background: linear-gradient(135deg, #e8f5e8 0%, #c8e6c9 100%);
}

.image-bottom .image-placeholder {
  background: linear-gradient(135deg, #fff3e0 0%, #ffe0b2 100%);
}

.image-back .image-placeholder {
  background: linear-gradient(135deg, #fbe9e7 0%, #ffccbc 100%);
}

.image-front .image-placeholder {
  background: linear-gradient(135deg, #e0f2f1 0%, #b2dfdb 100%);
}
</style>
