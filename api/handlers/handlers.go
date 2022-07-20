package handlers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/imayrus/url-shortener/database"
	"github.com/imayrus/url-shortener/models"
	"github.com/imayrus/url-shortener/random"
)

func Redirect(c *fiber.Ctx) error {

	Url := c.Params("url")
	url, err := FindByUrl(Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not find url in DB ",
		})
	}

	url.Clicked += 1
	err = UpdateUrl(url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error in updating",
		})
	}

	return c.Redirect(url.Url, fiber.StatusTemporaryRedirect)
}

func CreateUrl(c *fiber.Ctx) error {

	var body models.ShortUrl

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "can not parse json",
		})
	}

	if !govalidator.IsURL(body.Url) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid URL"})
	}

	Url := c.Params("url")
	resp, err := FindByUrl(Url)
	if err != nil {
		body.ShortUrl = random.RandomUrl(8)
		resp = models.ShortUrl{
			Url:      body.Url,
			ShortUrl: body.ShortUrl,
			Clicked:  body.Clicked,
		}
		err := Createurl(resp)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "could not create ShortUrl in db ",
			})
		}
		return c.Status(fiber.StatusOK).JSON(resp)

	}

	return c.Status(fiber.StatusOK).JSON(body)

}

func FindByUrl(Url string) (models.ShortUrl, error) {
	var findurl models.ShortUrl
	tx := database.GetDB().Where("Url = ?", Url).First(&findurl)
	return findurl, tx.Error
}

func UpdateUrl(Url models.ShortUrl) error {
	tx := database.GetDB().Save(&Url)
	return tx.Error
}

func Createurl(resp models.ShortUrl) error {
	tx := database.GetDB().Create(&resp)
	return tx.Error
}
