package providers

var ARCHIVE_PATH string = "/media/mvalgueiro/Arquivos/"

type Provider interface {
	sync() error
}
